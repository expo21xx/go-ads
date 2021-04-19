package ads

import (
	"encoding/binary"
)

// AMS TCP Header as defined by https://infosys.beckhoff.com/english.php?content=../content/1033/tc3_ads_intro/index.html
//
// +-----+----+---+---+---+---+
// | 0   | 1  | 2 | 3 | 4 | 5 |
// +-----+----+---+---+---+---+
// | reserved |     length    |
// +----------+---------------+
type amsTCPHeader struct {
	length uint32
}

func parseAMSTCPHeader(header [6]byte) amsTCPHeader {
	return amsTCPHeader{binary.LittleEndian.Uint32(header[2:])}
}

func newAMSTCPPacket(header *amsHeader, data []byte) []byte {
	var b = make([]byte, 38)

	if header.dataLength == 0 && len(data) != 0 {
		header.dataLength = uint32(len(data))
	}

	hb := header.bytes()

	binary.LittleEndian.PutUint32(b[2:6], uint32(len(hb)+len(data)))
	copy(b[6:38], hb)

	b = append(b, data...)

	return b
}
