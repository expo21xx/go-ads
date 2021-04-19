package ads

import (
	"encoding/binary"
)

// AMS Header as defined by https://infosys.beckhoff.com/english.php?content=../content/1033/tc3_ads_intro/index.html
//
// +------+-----+------+------+---+---+--------+-------+
// | 0    | 1   | 2    | 3    | 4 | 5 | 6      | 7     |
// +------+-----+------+------+---+---+--------+-------+
// |          AMSNetId Target         | AMSPort Target |
// +----------------------------------+----------------+
// |          AMSNetId Source         | AMSPort Source |
// +------------+-------------+-------+----------------+
// | Command Id | State Flags |       Data Length      |
// +------------+-------------+------------------------+
// |        Error Code        |        Invoke Id       |
// +--------------------------+------------------------+
// |                        Data                       |
// +---------------------------------------------------+
type amsHeader struct {
	targetNetID NetID
	targetPort  NetPort

	sourceNetID NetID
	sourcePort  NetPort

	commandID uint16

	stateFlags uint16

	dataLength uint32

	errorCode uint32

	invokeID uint32
}

func (h *amsHeader) isResponse() bool {
	return h.stateFlags == 0x0005
}

func parseAMSHeader(b [32]byte) *amsHeader {
	h := &amsHeader{}

	copy(h.targetNetID[:], b[0:6])
	h.targetPort = NetPort(binary.LittleEndian.Uint16(b[6:8]))

	copy(h.sourceNetID[:], b[8:14])
	h.sourcePort = NetPort(binary.LittleEndian.Uint16(b[14:16]))

	h.commandID = binary.LittleEndian.Uint16(b[16:18])
	h.stateFlags = binary.LittleEndian.Uint16(b[18:20])

	h.dataLength = binary.LittleEndian.Uint32(b[20:24])

	h.errorCode = binary.LittleEndian.Uint32(b[24:28])

	h.invokeID = binary.LittleEndian.Uint32(b[28:32])

	return h
}

func (h *amsHeader) bytes() []byte {
	var b [32]byte

	copy(b[0:6], h.targetNetID[:])
	binary.LittleEndian.PutUint16(b[6:8], uint16(h.targetPort))
	copy(b[8:14], h.sourceNetID[:])
	binary.LittleEndian.PutUint16(b[14:16], uint16(h.sourcePort))

	binary.LittleEndian.PutUint16(b[16:18], h.commandID)
	binary.LittleEndian.PutUint16(b[18:20], h.stateFlags)

	binary.LittleEndian.PutUint32(b[20:24], h.dataLength)
	binary.LittleEndian.PutUint32(b[24:28], h.errorCode)
	binary.LittleEndian.PutUint32(b[28:32], h.invokeID)

	return b[:]
}

func newAMSResponseHeader(reqHeader *amsHeader) *amsHeader {
	return &amsHeader{
		targetNetID: reqHeader.sourceNetID,
		targetPort:  reqHeader.sourcePort,
		sourceNetID: reqHeader.targetNetID,
		sourcePort:  reqHeader.targetPort,
		commandID:   reqHeader.commandID,
		errorCode:   ADSErrNoError,
		invokeID:    reqHeader.invokeID,
		stateFlags:  0x0005,
	}
}
