package ads

import (
	"encoding/binary"
	"errors"
	"fmt"
	"time"
)

const (
	// CommandInvalid .
	CommandInvalid = 0x0000

	// CommandADSReadDeviceInfo .
	CommandADSReadDeviceInfo = 0x0001

	// CommandADSRead .
	CommandADSRead = 0x0002

	// CommandADSWrite .
	CommandADSWrite = 0x0003

	// CommandADSReadState .
	CommandADSReadState = 0x0004

	// CommandADSWriteControl .
	CommandADSWriteControl = 0x0005

	// CommandADSAddDeviceNotification .
	CommandADSAddDeviceNotification = 0x0006

	// CommandADSDeletDeviceNotification .
	CommandADSDeletDeviceNotification = 0x0007

	// CommandADSDeviceNotification .
	CommandADSDeviceNotification = 0x0008

	// CommandADSReadWrite .
	CommandADSReadWrite = 0x0009
)

// CommandIDToString .
func CommandIDToString(cmd int) string {
	switch cmd {
	case CommandInvalid:
		return "invalid command (0x0000)"
	case CommandADSReadDeviceInfo:
		return "read device info command (0x0001)"
	case CommandADSRead:
		return "read command (0x0002)"
	case CommandADSWrite:
		return "write command (0x0003)"
	case CommandADSReadState:
		return "read state command (0x0004)"
	case CommandADSWriteControl:
		return "write control command (0x0005)"
	case CommandADSAddDeviceNotification:
		return "add device notification command (0x0006)"
	case CommandADSDeletDeviceNotification:
		return "delete device notification command (0x0007)"
	case CommandADSDeviceNotification:
		return "device notification command (0x0008)"
	case CommandADSReadWrite:
		return "read write command (0x0009)"
	}

	return fmt.Sprintf("unkown command id %v", cmd)
}

// Cmd .
type Cmd interface {
	Tag() int
	Bytes() []byte
	FromBytes([]byte) error
}

// ReadDeviceInfoCmdRequest see: https://infosys.beckhoff.com/english.php?content=../content/1033/tc3_ads_intro/index.html
// Reads the name and the version number of the ADS device.
type ReadDeviceInfoCmdRequest struct{}

// Tag .
func (r *ReadDeviceInfoCmdRequest) Tag() int {
	return CommandADSReadDeviceInfo
}

// Bytes .
func (r *ReadDeviceInfoCmdRequest) Bytes() []byte {
	return []byte{}
}

// FromBytes .
func (r *ReadDeviceInfoCmdRequest) FromBytes([]byte) error {
	return nil
}

// ReadDeviceInfoCmdResponse see: https://infosys.beckhoff.com/english.php?content=../content/1033/tc3_ads_intro/index.html
// Reads the name and the version number of the ADS device.
type ReadDeviceInfoCmdResponse struct {
	Result       uint32
	MajorVersion uint8
	MinorVersion uint8
	VersionBuild uint16
	DeviceName   string
}

func (r *ReadDeviceInfoCmdResponse) fromBytes(b []byte) error {
	if len(b) < 24 {
		return fmt.Errorf("too few bytes (%v) supplied. Minimum required %v", len(b), 24)
	}

	r.Result = binary.LittleEndian.Uint32(b[0:4])
	r.MajorVersion = b[4]
	r.MinorVersion = b[5]
	r.VersionBuild = binary.LittleEndian.Uint16(b[6:8])

	deviceName := make([]byte, 0, 16)
	for _, c := range b[8:24] {
		if c == 0 {
			break
		}

		deviceName = append(deviceName, c)
	}

	r.DeviceName = string(deviceName)

	return nil
}

// ReadCmdRequest see: https://infosys.beckhoff.com/english.php?content=../content/1033/tc3_ads_intro/index.html
// With ADS Read data can be read from an ADS device. The data are addressed by the Index Group and the Index Offset
type ReadCmdRequest struct {
	IndexGroup  uint32
	IndexOffset uint32
	Length      uint32
}

// Tag .
func (r *ReadCmdRequest) Tag() int {
	return CommandADSRead
}

// Bytes .
func (r *ReadCmdRequest) Bytes() []byte {
	b := [12]byte{}

	binary.LittleEndian.PutUint32(b[0:4], r.IndexGroup)
	binary.LittleEndian.PutUint32(b[4:8], r.IndexOffset)
	binary.LittleEndian.PutUint32(b[8:12], r.Length)

	return b[:]
}

// FromBytes .
func (r *ReadCmdRequest) FromBytes(b []byte) error {
	if len(b) < 12 {
		return fmt.Errorf("too few bytes (%v) supplied. Minimum required %v", len(b), 12)
	}

	r.IndexGroup = binary.LittleEndian.Uint32(b[0:4])
	r.IndexOffset = binary.LittleEndian.Uint32(b[4:8])
	r.Length = binary.LittleEndian.Uint32(b[8:12])

	return nil
}

// ReadCmdResponse see: https://infosys.beckhoff.com/english.php?content=../content/1033/tc3_ads_intro/index.html
// With ADS Read data can be read from an ADS device. The data are addressed by the Index Group and the Index Offset
type ReadCmdResponse struct {
	Result uint32
	Data   []byte
}

func (r *ReadCmdResponse) fromBytes(b []byte) error {
	if len(b) < 8 {
		return fmt.Errorf("too few bytes (%v) supplied. Minimum required %v", len(b), 8)
	}

	r.Result = binary.LittleEndian.Uint32(b[0:4])

	length := binary.LittleEndian.Uint32(b[4:8])
	r.Data = b[8 : 8+length]

	return nil
}

// WriteCmdRequest see: https://infosys.beckhoff.com/english.php?content=../content/1033/tc3_ads_intro/index.html&id
// With ADS Write data can be written to an ADS device. The data are addressed by the Index Group and the Index Offset
type WriteCmdRequest struct {
	IndexGroup  uint32
	IndexOffset uint32
	Data        []byte
}

// Tag .
func (r *WriteCmdRequest) Tag() int {
	return CommandADSWrite
}

// Bytes .
func (r *WriteCmdRequest) Bytes() []byte {
	length := uint32(len(r.Data))

	b := make([]byte, length+12)

	binary.LittleEndian.PutUint32(b[0:4], r.IndexGroup)
	binary.LittleEndian.PutUint32(b[4:8], r.IndexOffset)
	binary.LittleEndian.PutUint32(b[8:12], length)

	copy(b[12:], r.Data)

	return b
}

// FromBytes .
func (r *WriteCmdRequest) FromBytes(b []byte) error {
	if len(b) < 12 {
		return fmt.Errorf("too few bytes (%v) supplied. Minimum required %v", len(b), 12)
	}

	r.IndexGroup = binary.LittleEndian.Uint32(b[0:4])
	r.IndexOffset = binary.LittleEndian.Uint32(b[4:8])
	length := binary.LittleEndian.Uint32(b[8:12])

	r.Data = make([]byte, length)
	copy(r.Data, b[12:])

	return nil
}

// WriteCmdResponse see: https://infosys.beckhoff.com/english.php?content=../content/1033/tc3_ads_intro/index.html&id
// With ADS Write data can be written to an ADS device. The data are addressed by the Index Group and the Index Offset
type WriteCmdResponse struct {
	Result uint32
}

func (r *WriteCmdResponse) fromBytes(b []byte) error {
	if len(b) < 4 {
		return fmt.Errorf("too few bytes (%v) supplied. Minimum required %v", len(b), 4)
	}
	r.Result = binary.LittleEndian.Uint32(b[0:4])
	return nil
}

// ReadStateCmdRequest see: https://infosys.beckhoff.com/english.php?content=../content/1033/tc3_ads_intro/index.html&id=
// Reads the ADS status and the device status of an ADS device.
type ReadStateCmdRequest struct{}

// Tag .
func (r *ReadStateCmdRequest) Tag() int {
	return CommandADSReadState
}

// Bytes .
func (r *ReadStateCmdRequest) Bytes() []byte {
	return []byte{}
}

// FromBytes .
func (r *ReadStateCmdRequest) FromBytes(b []byte) error {
	return nil
}

// ReadStateCmdResponse see: https://infosys.beckhoff.com/english.php?content=../content/1033/tc3_ads_intro/index.html&id=
// Reads the ADS status and the device status of an ADS device.
type ReadStateCmdResponse struct {
	Result      uint32
	ADSState    uint16
	DeviceState uint16
}

func (r *ReadStateCmdResponse) fromBytes(b []byte) error {
	if len(b) < 8 {
		return fmt.Errorf("too few bytes (%v) supplied. Minimum required %v", len(b), 8)
	}

	r.Result = binary.LittleEndian.Uint32(b[0:4])
	r.ADSState = binary.LittleEndian.Uint16(b[4:6])
	r.DeviceState = binary.LittleEndian.Uint16(b[6:8])

	return nil
}

// FromBytes .
func (r *ReadStateCmdResponse) bytes() []byte {
	var b [8]byte

	binary.LittleEndian.PutUint32(b[0:4], r.Result)
	binary.LittleEndian.PutUint16(b[4:6], r.ADSState)
	binary.LittleEndian.PutUint16(b[6:8], r.DeviceState)

	return b[:]
}

// WriteControlCmdRequest see: https://infosys.beckhoff.com/english.php?content=../content/1033/tc3_ads_intro/index.html&id=
// Changes the ADS status and the device status of an ADS device. Additionally it is possible to send data to the ADS device to transfer further information.
// These data were not analyzed from the current ADS devices (PLC, NC, ...)
type WriteControlCmdRequest struct {
	ADSState    uint16
	DeviceState uint16
	Data        []byte
}

// Tag .
func (r *WriteControlCmdRequest) Tag() int {
	return CommandADSWriteControl
}

// FromBytes .
func (r *WriteControlCmdRequest) FromBytes(b []byte) error {
	if len(b) < 8 {
		return fmt.Errorf("too few bytes (%v) supplied. Minimum required %v", len(b), 8)
	}

	r.ADSState = binary.LittleEndian.Uint16(b[0:2])
	r.DeviceState = binary.LittleEndian.Uint16(b[2:4])

	length := binary.LittleEndian.Uint32(b[4:8])

	r.Data = make([]byte, length)
	copy(r.Data, b[8:])

	return nil
}

// Bytes .
func (r *WriteControlCmdRequest) Bytes() []byte {
	length := uint32(len(r.Data))

	b := make([]byte, length+12)

	binary.LittleEndian.PutUint16(b[0:2], r.ADSState)
	binary.LittleEndian.PutUint16(b[2:4], r.DeviceState)
	binary.LittleEndian.PutUint32(b[4:8], length)

	copy(b[8:], r.Data)

	return b
}

// WriteControlCmdResponse see: https://infosys.beckhoff.com/english.php?content=../content/1033/tc3_ads_intro/index.html&id=
// Changes the ADS status and the device status of an ADS device. Additionally it is possible to send data to the ADS device to transfer further information.
// These data were not analyzed from the current ADS devices (PLC, NC, ...)
type WriteControlCmdResponse struct {
	Result uint32
}

func (r *WriteControlCmdResponse) fromBytes(b []byte) error {
	if len(b) < 4 {
		return fmt.Errorf("too few bytes (%v) supplied. Minimum required %v", len(b), 4)
	}

	r.Result = binary.LittleEndian.Uint32(b[0:4])
	return nil
}

// AddDeviceNotificationCmdRequest see: https://infosys.beckhoff.com/english.php?content=../content/1033/tc3_ads_intro/index.html&id=
// A notification is created in an ADS device.
// Note: We recommend to announce not more than 550 notifications per device. Otherwise increase the payload by working with structures or use sum commands.
type AddDeviceNotificationCmdRequest struct {
	// Index Group of the data, which should be sent per notification.
	IndexGroup uint32
	// Index Offset of the data, which should be sent per notification.
	IndexOffset uint32
	// Length of data in bytes, which should be sent per notification.
	Length           uint32
	TransmissionMode uint32
	// At the latest after this time, the ADS Device Notification is called. The unit is 1ms.
	MaxDelay uint32
	// The ADS server checks if the value changes in this time slice. The unit is 1ms.
	CycleTime uint32
}

// Tag .
func (r *AddDeviceNotificationCmdRequest) Tag() int {
	return CommandADSAddDeviceNotification
}

// FromBytes .
func (r *AddDeviceNotificationCmdRequest) FromBytes(b []byte) error {
	if len(b) < 32 {
		return fmt.Errorf("too few bytes (%v) supplied. Minimum required %v", len(b), 32)
	}

	r.IndexGroup = binary.LittleEndian.Uint32(b[0:4])
	r.IndexOffset = binary.LittleEndian.Uint32(b[4:8])
	r.Length = binary.LittleEndian.Uint32(b[8:12])
	r.TransmissionMode = binary.LittleEndian.Uint32(b[12:16])
	r.MaxDelay = binary.LittleEndian.Uint32(b[16:20])
	r.CycleTime = binary.LittleEndian.Uint32(b[20:24])

	return nil
}

// Bytes .
func (r *AddDeviceNotificationCmdRequest) Bytes() []byte {
	b := [32]byte{}

	binary.LittleEndian.PutUint32(b[0:4], r.IndexGroup)
	binary.LittleEndian.PutUint32(b[4:8], r.IndexOffset)
	binary.LittleEndian.PutUint32(b[8:12], r.Length)
	binary.LittleEndian.PutUint32(b[12:16], r.TransmissionMode)
	binary.LittleEndian.PutUint32(b[16:20], r.MaxDelay)
	binary.LittleEndian.PutUint32(b[20:24], r.CycleTime)

	return b[:]
}

// AddDeviceNotificationCmdResponse see: https://infosys.beckhoff.com/english.php?content=../content/1033/tc3_ads_intro/index.html&id=
// A notification is created in an ADS device.
// Note: We recommend to announce not more than 550 notifications per device. Otherwise increase the payload by working with structures or use sum commands.
type AddDeviceNotificationCmdResponse struct {
	Result             uint32
	NotificationHandle uint32
}

func (r *AddDeviceNotificationCmdResponse) fromBytes(b []byte) error {
	if len(b) < 8 {
		return fmt.Errorf("too few bytes (%v) supplied. Minimum required %v", len(b), 8)
	}

	r.Result = binary.LittleEndian.Uint32(b[0:4])
	r.NotificationHandle = binary.LittleEndian.Uint32(b[4:8])
	return nil
}

// DeleteDeviceNotificationCmdRequest see: https://infosys.beckhoff.com/english.php?content=../content/1033/tc3_ads_intro/index.html&id=
// One before defined notification is deleted in an ADS device.
type DeleteDeviceNotificationCmdRequest struct {
	NotificationHandle uint32
}

// Tag .
func (r *DeleteDeviceNotificationCmdRequest) Tag() int {
	return CommandADSDeletDeviceNotification
}

// FromBytes .
func (r *DeleteDeviceNotificationCmdRequest) FromBytes(b []byte) error {
	if len(b) < 4 {
		return fmt.Errorf("too few bytes (%v) supplied. Minimum required %v", len(b), 4)
	}

	r.NotificationHandle = binary.LittleEndian.Uint32(b[0:4])

	return nil
}

// Bytes .
func (r *DeleteDeviceNotificationCmdRequest) Bytes() []byte {
	b := [4]byte{}
	binary.LittleEndian.PutUint32(b[0:4], r.NotificationHandle)
	return b[:]
}

// DeleteDeviceNotificationCmdResponse see: https://infosys.beckhoff.com/english.php?content=../content/1033/tc3_ads_intro/index.html&id=
// One before defined notification is deleted in an ADS device.
type DeleteDeviceNotificationCmdResponse struct {
	Result uint32
}

func (r *DeleteDeviceNotificationCmdResponse) fromBytes(b []byte) error {
	if len(b) < 4 {
		return fmt.Errorf("too few bytes (%v) supplied. Minimum required %v", len(b), 4)
	}
	r.Result = binary.LittleEndian.Uint32(b[0:4])
	return nil
}

// DeviceNotificationCmdRequest see: https://infosys.beckhoff.com/english.php?content=../content/1033/tc3_ads_intro/index.html
// Data will carry forward independently from an ADS device to a Client.
// The data which are transferred at the Device Notification are multiple nested into one another. The Notification Stream contains
// an array with elements of type AdsStampHeader. This array again contains elements of type AdsNotificationSample.
type DeviceNotificationCmdRequest struct {
	Stamps []StampHeader
}

// Tag .
func (r *DeviceNotificationCmdRequest) Tag() int {
	return CommandADSDeviceNotification
}

// FromBytes .
func (r *DeviceNotificationCmdRequest) FromBytes(b []byte) error {
	numStamps := binary.LittleEndian.Uint32(b[4:8])
	r.Stamps = make([]StampHeader, numStamps)

	offset := 8
	for i := range r.Stamps {
		advance, err := r.Stamps[i].fromBytes(b[offset:])
		if err != nil {
			return err
		}

		offset = offset + advance
	}

	return nil
}

// Bytes .
func (r *DeviceNotificationCmdRequest) Bytes() []byte {
	b := make([]byte, 8)

	for i := range r.Stamps {
		b = append(b, r.Stamps[i].bytes()...)
	}

	binary.LittleEndian.PutUint32(b[0:4], uint32(len(b)-12))
	binary.LittleEndian.PutUint32(b[4:8], uint32(len(r.Stamps)))

	return b
}

// StampHeader see: https://infosys.beckhoff.com/english.php?content=../content/1033/tc3_ads_intro/index.html
type StampHeader struct {
	Timestamp time.Time
	Samples   []Sample
}

func (s *StampHeader) fromBytes(b []byte) (int, error) {
	if len(b) < 12 {
		return 0, fmt.Errorf("too few bytes (%v) supplied. Minimum required %v", len(b), 12)
	}

	// The timestamp is coded after the Windows FILETIME format.
	//	I.e. the value contains the number of the nano seconds, which passed since 1.1.1601.
	//	In addition, the local time change is not considered. Thus the time stamp is present
	// as universal Coordinated time (UTC).
	ft := &filetime{
		lowDateTime:  binary.LittleEndian.Uint32(b[0:4]),
		highDateTime: binary.LittleEndian.Uint32(b[4:8]),
	}

	s.Timestamp = time.Unix(0, ft.nanoseconds()).UTC()

	numSamples := binary.LittleEndian.Uint32(b[8:12])

	s.Samples = make([]Sample, numSamples)

	offset := 12
	for i := range s.Samples {
		advance, err := s.Samples[i].fromBytes(b[offset:])
		if err != nil {
			return 0, err
		}

		offset = offset + advance
	}

	return offset, nil
}

func (s *StampHeader) bytes() []byte {
	b := make([]byte, 12)

	ft := nsecToFiletime(int64(s.Timestamp.UTC().UnixNano()))

	binary.LittleEndian.PutUint32(b[0:4], ft.lowDateTime)
	binary.LittleEndian.PutUint32(b[4:8], ft.highDateTime)
	binary.LittleEndian.PutUint32(b[8:12], uint32(len(s.Samples)))

	for i := range s.Samples {
		b = append(b, s.Samples[i].bytes()...)
	}

	return b
}

// Sample see: https://infosys.beckhoff.com/english.php?content=../content/1033/tc3_ads_intro/index.html
type Sample struct {
	NotificationHandle uint32
	Data               []byte
}

func (s *Sample) fromBytes(b []byte) (int, error) {
	if len(b) < 8 {
		return 0, fmt.Errorf("too few bytes (%v) supplied. Minimum required %v", len(b), 8)
	}

	s.NotificationHandle = binary.LittleEndian.Uint32(b[0:4])

	length := binary.LittleEndian.Uint32(b[4:8])
	s.Data = make([]byte, length)

	copied := copy(s.Data, b[8:])

	return copied + 8, nil
}

func (s *Sample) bytes() []byte {
	length := uint32(len(s.Data) + 8)
	b := make([]byte, length)

	binary.LittleEndian.PutUint32(b[0:4], s.NotificationHandle)
	binary.LittleEndian.PutUint32(b[4:8], uint32(len(s.Data)))

	copy(b[8:], s.Data)

	return b
}

// ReadWriteCmdRequest see: https://infosys.beckhoff.com/english.php?content=../content/1033/tc3_ads_intro/index.html
// With ADS ReadWrite data will be written to an ADS device. Additionally, data can be read from the ADS device.
// The data which can be read are addressed by the Index Group and the Index Offset
type ReadWriteCmdRequest struct {
	IndexGroup  uint32
	IndexOffset uint32
	ReadLength  uint32
	Data        []byte
}

// Tag .
func (r *ReadWriteCmdRequest) Tag() int {
	return CommandADSReadWrite
}

// FromBytes .
func (r *ReadWriteCmdRequest) FromBytes(b []byte) error {
	if len(b) < 16 {
		return fmt.Errorf("too few bytes (%v) supplied. Minimum required %v", len(b), 16)
	}

	r.IndexGroup = binary.LittleEndian.Uint32(b[0:4])
	r.IndexOffset = binary.LittleEndian.Uint32(b[4:8])
	r.ReadLength = binary.LittleEndian.Uint32(b[8:12])

	length := binary.LittleEndian.Uint32(b[12:16])
	r.Data = make([]byte, length)
	copy(r.Data, b[16:16+length])

	return nil
}

// Bytes .
func (r *ReadWriteCmdRequest) Bytes() []byte {
	length := uint32(len(r.Data))

	b := make([]byte, length+16)

	binary.LittleEndian.PutUint32(b[0:4], r.IndexGroup)
	binary.LittleEndian.PutUint32(b[4:8], r.IndexOffset)
	binary.LittleEndian.PutUint32(b[8:12], r.ReadLength)
	binary.LittleEndian.PutUint32(b[12:16], length)

	copy(b[16:], r.Data)

	return b
}

// ReadWriteCmdResponse see: https://infosys.beckhoff.com/english.php?content=../content/1033/tc3_ads_intro/index.html
// With ADS ReadWrite data will be written to an ADS device. Additionally, data can be read from the ADS device.
// The data which can be read are addressed by the Index Group and the Index Offset
type ReadWriteCmdResponse struct {
	Result uint32
	Data   []byte
}

func (r *ReadWriteCmdResponse) fromBytes(b []byte) error {
	if len(b) < 4 {
		return errors.New("too few bytes supplied")
	}

	r.Result = binary.LittleEndian.Uint32(b[0:4])

	length := binary.LittleEndian.Uint32(b[4:8])
	r.Data = make([]byte, length)

	copy(r.Data, b[8:8+length])

	return nil
}
