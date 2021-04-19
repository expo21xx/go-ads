package ads

import (
	"bytes"
	"context"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"sync"
	"sync/atomic"
	"time"
)

// Client is the primary way to communicate with an ADS/AMS system.
type Client struct {
	mu sync.Mutex

	conn net.Conn

	ip            net.IP
	targetNetID   NetID
	targetNetPort NetPort

	sourceNetID   NetID
	sourceNetPort NetPort

	receive  chan []byte
	transmit chan transmission
	shutdown chan bool
	closing  bool

	invokeID uint32

	responseChans map[uint32]chan response

	symbols map[string]Symbol

	notificationHandlers           map[uint32]func(interface{})
	notificationHandlersToDataType map[uint32]string

	loadSymbolsOnStart bool
	monitorSymbols     bool
	monitorSymbolsStop func() error

	err error
}

type response struct {
	err    error
	header *amsHeader
	data   []byte
}

type transmission struct {
	responseChan chan response
	data         []byte
}

// NewClient .
func NewClient(ip, netid string, port int, opts ...Option) (*Client, error) {
	var targetNetID NetID
	var err error

	if netid != "" {
		targetNetID, err = ParseNetIDFromString(netid)
		if err != nil {
			return nil, err
		}
	}

	c := &Client{
		ip:                             net.ParseIP(ip),
		targetNetPort:                  NetPort(port),
		targetNetID:                    targetNetID,
		loadSymbolsOnStart:             false,
		monitorSymbols:                 false,
		sourceNetPort:                  800,
		symbols:                        map[string]Symbol{},
		responseChans:                  map[uint32]chan response{},
		notificationHandlers:           map[uint32]func(interface{}){},
		notificationHandlersToDataType: map[uint32]string{},
	}

	for _, opt := range opts {
		opt.apply(c)
	}

	return c, nil
}

// Connect .
func (c *Client) Connect(ctx context.Context) error {
	var d net.Dialer

	conn, err := d.DialContext(ctx, "tcp", fmt.Sprintf("%v:%v", c.ip.String(), ADSTCPServerPort))
	if err != nil {
		return err
	}

	c.conn = conn

	if c.sourceNetID.String() == "0.0.0.0.0.0" {
		host, _, _ := net.SplitHostPort(conn.LocalAddr().String())
		sourceNetID, err := ParseNetIDFromString(host + ".1.1")
		if err != nil {
			return err
		}
		c.sourceNetID = sourceNetID
	}

	c.receive = make(chan []byte)
	c.transmit = make(chan transmission)
	c.shutdown = make(chan bool)

	go c.receiver()
	go c.transmitter()

	if c.loadSymbolsOnStart {
		err = c.FetchSymbols(ctx)
		if err != nil {
			c.Close(ctx)
			return err
		}
	}

	if c.monitorSymbols {
		c.monitorSymbolsStop, err = c.setupMonitorSymbols(context.Background())
		if err != nil {
			c.Close(ctx)
			return err
		}
	}

	return nil
}

// Close .
func (c *Client) Close(ctx context.Context) error {
	c.mu.Lock()
	monitorSymbolsStop := c.monitorSymbolsStop

	if c.conn == nil {
		return nil
	}

	c.closing = true

	c.mu.Unlock()

	if monitorSymbolsStop != nil {
		monitorSymbolsStop()
	}

	handles := make([]uint32, len(c.notificationHandlers))
	i := 0
	c.mu.Lock()
	for handle := range c.notificationHandlers {
		handles[i] = handle
		i++
	}
	c.mu.Unlock()

	c.notificationHandlers = map[uint32]func(interface{}){}

	for _, h := range handles {
		err := c.DeleteDeviceNotification(ctx, h)
		if err != nil {
			return err
		}
	}

	err := c.conn.Close()
	if err != nil {
		return err
	}

	c.conn = nil

	close(c.receive)
	close(c.transmit)
	close(c.shutdown)

	c.mu.Lock()
	c.closing = false
	c.mu.Unlock()

	return nil
}

// AddSymbol .
func (c *Client) AddSymbol(sym Symbol) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.symbols[sym.Name] = sym
}

// GetSymbol .
func (c *Client) GetSymbol(name string) (Symbol, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	s, ok := c.symbols[name]
	return s, ok
}

// ReadDeviceInfo .
func (c *Client) ReadDeviceInfo(ctx context.Context) (*ReadDeviceInfoCmdResponse, error) {
	cmd := &ReadDeviceInfoCmdRequest{}

	data, err := c.sendRequest(ctx, cmd)
	if err != nil {
		return nil, err
	}

	res := &ReadDeviceInfoCmdResponse{}
	err = res.fromBytes(data)
	if err != nil {
		return nil, err
	}

	if res.Result != ADSErrNoError {
		return nil, fmt.Errorf("read device info returned error code: %v (0x%x)", ErrCodeToString(res.Result), res.Result)
	}

	return res, nil
}

// Read .
func (c *Client) Read(ctx context.Context, indexGroup uint32, indexOffset uint32, length int) ([]byte, error) {
	cmd := &ReadCmdRequest{
		IndexGroup:  indexGroup,
		IndexOffset: indexOffset,
		Length:      uint32(length),
	}

	data, err := c.sendRequest(ctx, cmd)
	if err != nil {
		return nil, err
	}

	res := &ReadCmdResponse{}
	err = res.fromBytes(data)
	if err != nil {
		return nil, err
	}

	if res.Result != ADSErrNoError {
		return nil, fmt.Errorf("read returned error code: %v (0x%x)", ErrCodeToString(res.Result), res.Result)
	}

	return res.Data, nil
}

// ReadByName .
func (c *Client) ReadByName(ctx context.Context, name string) (interface{}, error) {
	symbol, ok := c.GetSymbol(name)
	if !ok {
		return nil, fmt.Errorf("unkown symbol %v", name)
	}

	val, err := c.Read(ctx, symbol.IndexGroup, symbol.IndexOffset, int(symbol.Size))
	if err != nil {
		return nil, err
	}

	return adsBytesToGoValue(symbol.Type, val)
}

// ReadByNameBytes .
func (c *Client) ReadByNameBytes(ctx context.Context, name string, size int) ([]byte, error) {
	handle, err := c.getSymHandleByName(ctx, name)
	if err != nil {
		return nil, err
	}

	val, err := c.Read(ctx, handle, 0, size)
	if err != nil {
		return nil, err
	}

	return val, nil
}

// Write .
func (c *Client) Write(ctx context.Context, indexGroup uint32, indexOffset uint32, data []byte) error {
	cmd := &WriteCmdRequest{
		IndexGroup:  indexGroup,
		IndexOffset: indexOffset,
		Data:        data,
	}

	data, err := c.sendRequest(ctx, cmd)
	if err != nil {
		return err
	}

	res := &WriteCmdResponse{}
	err = res.fromBytes(data)
	if err != nil {
		return err
	}

	if res.Result != ADSErrNoError {
		return fmt.Errorf("write returned error code: %v (0x%x)", ErrCodeToString(res.Result), res.Result)
	}

	return nil
}

// WriteByName .
func (c *Client) WriteByName(ctx context.Context, name string, data []byte) error {
	symbol, ok := c.GetSymbol(name)
	if !ok {
		return fmt.Errorf("unkown symbol %v", name)
	}

	return c.Write(ctx, symbol.IndexGroup, symbol.IndexOffset, data)
}

// ReadState .
func (c *Client) ReadState(ctx context.Context) (adsState uint16, deviceState uint16, err error) {
	cmd := &ReadStateCmdRequest{}

	data, err := c.sendRequest(ctx, cmd)
	if err != nil {
		return 0, 0, err
	}

	res := &ReadStateCmdResponse{}
	err = res.fromBytes(data)
	if err != nil {
		return 0, 0, err
	}

	return res.ADSState, res.DeviceState, nil
}

// WriteControl .
func (c *Client) WriteControl(ctx context.Context, adsState uint16, deviceState uint16, data []byte) error {
	cmd := &WriteControlCmdRequest{
		ADSState:    adsState,
		DeviceState: deviceState,
		Data:        data,
	}

	data, err := c.sendRequest(ctx, cmd)
	if err != nil {
		return err
	}

	res := &WriteControlCmdResponse{}
	err = res.fromBytes(data)
	if err != nil {
		return err
	}

	if res.Result != ADSErrNoError {
		return fmt.Errorf("write control returned error code: %v (0x%x)", ErrCodeToString(res.Result), res.Result)
	}

	return nil
}

// AddDeviceNotification .
func (c *Client) AddDeviceNotification(ctx context.Context, req *AddDeviceNotificationCmdRequest) (uint32, error) {
	data, err := c.sendRequest(ctx, req)
	if err != nil {
		return 0, err
	}

	res := &AddDeviceNotificationCmdResponse{}
	err = res.fromBytes(data)
	if err != nil {
		return 0, err
	}

	if res.Result != ADSErrNoError {
		return 0, fmt.Errorf("add device notification returned error code: %v (0x%x)", ErrCodeToString(res.Result), res.Result)
	}

	return res.NotificationHandle, nil
}

// DeleteDeviceNotification .
func (c *Client) DeleteDeviceNotification(ctx context.Context, handle uint32) error {
	cmd := &DeleteDeviceNotificationCmdRequest{
		NotificationHandle: handle,
	}

	data, err := c.sendRequest(ctx, cmd)
	if err != nil {
		return err
	}

	res := &DeleteDeviceNotificationCmdResponse{}
	err = res.fromBytes(data)
	if err != nil {
		return err
	}

	if res.Result != ADSErrNoError {
		return fmt.Errorf("delete device notification returned error code: %v (0x%x)", ErrCodeToString(res.Result), res.Result)
	}

	return nil
}

// SendDeviceNotification .
func (c *Client) SendDeviceNotification(ctx context.Context, n *DeviceNotificationCmdRequest) error {
	_, err := c.sendRequest(ctx, n)
	if err != nil {
		return err
	}

	return nil
}

// ReadWrite .
func (c *Client) ReadWrite(ctx context.Context, indexGroup uint32, indexOffset uint32, data []byte) ([]byte, error) {
	cmd := &ReadWriteCmdRequest{
		IndexGroup:  indexGroup,
		IndexOffset: indexOffset,
		ReadLength:  uint32(len(data)),
		Data:        data,
	}

	data, err := c.sendRequest(ctx, cmd)
	if err != nil {
		return nil, err
	}

	res := &ReadWriteCmdResponse{}
	err = res.fromBytes(data)
	if err != nil {
		return nil, err
	}

	if res.Result != ADSErrNoError {
		return nil, fmt.Errorf("read write returned error code: %v (0x%x)", ErrCodeToString(res.Result), res.Result)
	}

	return res.Data, nil
}

// ReadWriteByName .
func (c *Client) ReadWriteByName(ctx context.Context, name string, data []byte) (interface{}, error) {
	symbol, ok := c.GetSymbol(name)
	if !ok {
		return nil, fmt.Errorf("unkown symbol %v", name)
	}

	val, err := c.ReadWrite(ctx, symbol.IndexGroup, symbol.IndexOffset, data)
	if err != nil {
		return nil, err
	}

	return adsBytesToGoValue(symbol.Type, val)
}

// AddDeviceNotificationHandler .
func (c *Client) AddDeviceNotificationHandler(ctx context.Context, req *AddDeviceNotificationCmdRequest, dt string, cb func(interface{})) (func() error, error) {
	handle, err := c.AddDeviceNotification(ctx, req)
	if err != nil {
		return func() error { return nil }, err
	}

	c.mu.Lock()
	if dt != "" {
		c.notificationHandlersToDataType[handle] = dt
	}

	c.notificationHandlers[handle] = cb
	c.mu.Unlock()

	return func() error {
		c.mu.Lock()
		delete(c.notificationHandlersToDataType, handle)
		if _, ok := c.notificationHandlers[handle]; !ok {
			c.mu.Unlock()
			return nil
		}
		delete(c.notificationHandlers, handle)
		c.mu.Unlock()

		return c.DeleteDeviceNotification(ctx, handle)
	}, nil
}

// DeviceNotificationOpts .
type DeviceNotificationOpts struct {
	TransmissionMode uint32
	// At the latest after this time, the ADS Device Notification is called. The unit is 1ms.
	MaxDelay uint32
	// The ADS server checks if the value changes in this time slice. The unit is 1ms.
	CycleTime uint32
}

// AddDeviceNotificationHandlerByName .
func (c *Client) AddDeviceNotificationHandlerByName(ctx context.Context, name string, opts *DeviceNotificationOpts, cb func(interface{})) (func() error, error) {
	if opts == nil {
		opts = &DeviceNotificationOpts{
			TransmissionMode: ADSTransServerOnCha,
			MaxDelay:         100,
			CycleTime:        100,
		}
	}

	symbol, ok := c.GetSymbol(name)
	if !ok {
		return func() error { return nil }, fmt.Errorf("unkown symbol %v", name)
	}

	handle, err := c.AddDeviceNotification(ctx, &AddDeviceNotificationCmdRequest{
		IndexGroup:       symbol.IndexGroup,
		IndexOffset:      symbol.IndexOffset,
		Length:           symbol.Size,
		TransmissionMode: opts.TransmissionMode,
		MaxDelay:         opts.MaxDelay,
		CycleTime:        opts.CycleTime,
	})
	if err != nil {
		return func() error { return nil }, err
	}

	c.mu.Lock()
	c.notificationHandlersToDataType[handle] = symbol.Type
	c.notificationHandlers[handle] = cb
	c.mu.Unlock()

	return func() error {
		c.mu.Lock()
		delete(c.notificationHandlersToDataType, handle)
		if _, ok := c.notificationHandlers[handle]; !ok {
			c.mu.Unlock()
			return nil
		}

		delete(c.notificationHandlers, handle)
		c.mu.Unlock()

		return c.DeleteDeviceNotification(ctx, handle)
	}, nil
}

// FetchSymbols .
func (c *Client) FetchSymbols(ctx context.Context) error {
	infoData, err := c.Read(ctx, ADSIndexGroupSymUploadInfo2, 0, 24)
	if err != nil {
		return err
	}

	var symInfo adsSymbolInfo
	buf := bytes.NewBuffer(infoData)
	binary.Read(buf, binary.LittleEndian, &symInfo)

	symbolData, err := c.Read(ctx, ADSIndexGroupSymUpload, 0, int(symInfo.DataTypeLength))
	if err != nil {
		return err
	}

	offset := 0
	for offset < len(symbolData) {
		var symbol Symbol
		length := binary.LittleEndian.Uint32(symbolData[offset : offset+4])

		symbol.IndexGroup = binary.LittleEndian.Uint32(symbolData[offset+4 : offset+8])
		symbol.IndexOffset = binary.LittleEndian.Uint32(symbolData[offset+8 : offset+12])
		symbol.Size = binary.LittleEndian.Uint32(symbolData[offset+12 : offset+16])
		// symbol.Type = binary.LittleEndian.Uint32(symbolData[offset+16 : offset+20])
		symbol.Flags = binary.LittleEndian.Uint32(symbolData[offset+20 : offset+24])

		nameLen := int(binary.LittleEndian.Uint16(symbolData[offset+24 : offset+26]))
		typeLen := int(binary.LittleEndian.Uint16(symbolData[offset+26 : offset+28]))
		commentLen := int(binary.LittleEndian.Uint16(symbolData[offset+28 : offset+30]))

		relOffset := offset + 30

		symbol.Name = string(symbolData[relOffset : relOffset+nameLen])

		relOffset = relOffset + nameLen + 1

		symbol.Type = string(symbolData[relOffset : relOffset+typeLen])

		relOffset = relOffset + typeLen + 2

		if commentLen != 0 {
			symbol.Comment = string(symbolData[relOffset : relOffset+commentLen-1])
		}

		c.AddSymbol(symbol)

		offset = offset + int(length)
	}

	return nil
}

func (c *Client) sendRequest(ctx context.Context, cmd Cmd) ([]byte, error) {
	header := &amsHeader{
		commandID:   uint16(cmd.Tag()),
		targetNetID: c.targetNetID,
		targetPort:  c.targetNetPort,
		sourceNetID: c.sourceNetID,
		sourcePort:  c.sourceNetPort,
		invokeID:    c.newInvokeID(),
		stateFlags:  0x0004,
	}

	data := cmd.Bytes()

	resChan := make(chan response)
	c.mu.Lock()
	c.responseChans[header.invokeID] = resChan
	c.mu.Unlock()
	defer func() {
		c.mu.Lock()
		close(c.responseChans[header.invokeID])
		delete(c.responseChans, header.invokeID)
		c.mu.Unlock()
	}()

	c.transmit <- transmission{responseChan: resChan, data: newAMSTCPPacket(header, data)}

	var res response
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	case res = <-resChan:
	case <-c.shutdown:
		return nil, errors.New("closing connection")
	}

	return res.data, res.err
}

func (c *Client) newInvokeID() uint32 { /*{{{*/
	return atomic.AddUint32(&c.invokeID, 1)
}

func (c *Client) setError(err error) {
	fmt.Println(err)
	c.mu.Lock()
	defer c.mu.Unlock()
	c.err = err
}

func (c *Client) receiver() {
	var buf bytes.Buffer
	b := make([]byte, 2048)
	var headerBytes [32]byte
	partial := false
	var header *amsHeader

	for {
		n, err := c.conn.Read(b)
		if err == io.EOF {
			return
		}

		if err != nil {
			c.mu.Lock()
			closing := c.closing
			c.mu.Unlock()

			if closing {
				return
			}

			c.setError(err)
			return
		}

		if n > 0 {
			buf.Write(b[:n])
		}

		// impartial message
		if buf.Len() < 38 {
			partial = true
			continue
		}

		if !partial {
			buf.Next(6)
			hn, err := buf.Read(headerBytes[:])
			if err != nil {
				log.Fatal(err)
			}

			if hn != 32 {
				log.Fatalf("not enough bytes read: %v", hn)
			}

			header = parseAMSHeader(headerBytes)
		}

		// impartial message
		if header.dataLength > uint32(buf.Len()) {
			partial = true
			continue
		}

		data := buf.Next(int(header.dataLength))
		dataCopy := make([]byte, len(data))
		// copy the bytes becasue the returned slice is only valid until the next read. This can be an issue when using notifications from multiple goroutines.
		copy(dataCopy[:], data)
		if header.isResponse() {
			c.handleResponse(header, dataCopy)
		} else {
			c.handleCommand(header, dataCopy)
		}

		buf.Reset()
		partial = false
		header = nil
	}
}

func (c *Client) handleResponse(header *amsHeader, data []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	resChan, ok := c.responseChans[header.invokeID]
	if !ok {
		err := fmt.Errorf("no response channel for invoke id: %v (command id %x)", header.invokeID, header.commandID)
		c.setError(err)
		return
	}

	if header.errorCode != ADSErrNoError {
		err := fmt.Errorf("received error: %v (0x%x)", ErrCodeToString(header.errorCode), header.errorCode)
		resChan <- response{header: header, err: err}
		return
	}

	resChan <- response{header: header, data: data}
}

func (c *Client) handleCommand(header *amsHeader, data []byte) {
	if header.errorCode != ADSErrNoError {
		err := fmt.Errorf("received error: %v (0x%x)", ErrCodeToString(header.errorCode), header.errorCode)
		c.setError(err)
		return
	}

	switch header.commandID {
	case CommandADSReadState:
		cmd := &ReadStateCmdRequest{}
		err := cmd.FromBytes(data)
		if err != nil {
			c.setError(err)
			return
		}

		c.readStateCommand(header, cmd)
	case CommandADSDeviceNotification:
		cmd := &DeviceNotificationCmdRequest{}
		err := cmd.FromBytes(data)
		if err != nil {
			c.setError(err)
			return
		}

		c.deviceNotificationCmd(header, cmd)
	default:
		c.setError(fmt.Errorf("unkown command ID %v", header.commandID))
	}

}

func (c *Client) deviceNotificationCmd(header *amsHeader, cmd *DeviceNotificationCmdRequest) {
	for _, stmp := range cmd.Stamps {
		for _, s := range stmp.Samples {
			c.mu.Lock()
			handler, ok := c.notificationHandlers[s.NotificationHandle]
			c.mu.Unlock()
			if !ok {
				c.setError(fmt.Errorf("no handlers for notification %v", s.NotificationHandle))
				continue
			}

			var val interface{} = s.Data
			var err error
			c.mu.Lock()
			if typ, ok := c.notificationHandlersToDataType[s.NotificationHandle]; ok {
				val, err = adsBytesToGoValue(typ, s.Data)
				if err != nil {
					c.setError(fmt.Errorf("error while converting notifation (%v) value to go value", s.NotificationHandle))
				}
			}
			c.mu.Unlock()

			handler(val)
		}
	}
}

func (c *Client) readStateCommand(header *amsHeader, cmd *ReadStateCmdRequest) {
	res := &ReadStateCmdResponse{
		Result:      ADSErrNoError,
		ADSState:    ADSStateRun,
		DeviceState: 1,
	}

	data := res.bytes()

	retHeader := newAMSResponseHeader(header)

	retHeader.dataLength = uint32(len(data))

	c.transmit <- transmission{data: newAMSTCPPacket(retHeader, data)}
}

func (c *Client) transmitter() {
	for {
		select {
		case t := <-c.transmit:
			var n int
			var err error
			for n < len(t.data) {
				n, err = c.conn.Write(t.data)
				if err != nil {
					if t.responseChan != nil {
						t.responseChan <- response{err: fmt.Errorf("error while writing: %w", err)}
					} else {
						c.setError(fmt.Errorf("error while writing: %w", err))
					}
				}
			}
		case <-c.shutdown:
			return
		}
	}
}

func (c *Client) setupMonitorSymbols(ctx context.Context) (func() error, error) {
	return c.AddDeviceNotificationHandler(ctx, &AddDeviceNotificationCmdRequest{
		IndexGroup:       ADSIndexGroupSymVersion,
		IndexOffset:      0,
		Length:           1,
		TransmissionMode: ADSTransServerOnCha,
		MaxDelay:         100,
		CycleTime:        1000,
	}, "", func(data interface{}) {
		go func() {
			ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
			defer cancel()
			c.FetchSymbols(ctx)
		}()
	})
}

func (c *Client) getSymHandleByName(ctx context.Context, name string) (uint32, error) {
	cmd := &ReadWriteCmdRequest{
		IndexGroup:  ADSIndexGroupSymHndByName,
		IndexOffset: 0,
		Data:        []byte(name),
		ReadLength:  4,
	}

	data, err := c.sendRequest(ctx, cmd)
	if err != nil {
		return 0, err
	}

	res := &ReadWriteCmdResponse{}
	err = res.fromBytes(data)
	if err != nil {
		return 0, err
	}
	if res.Result != ADSErrNoError {
		return 0, fmt.Errorf("getSymHandleByName (%v) returned error code: %v (0x%x)", name, ErrCodeToString(res.Result), res.Result)
	}

	return binary.LittleEndian.Uint32(res.Data), nil
}
