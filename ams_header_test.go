package ads

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParseAMSHeader(t *testing.T) {
	b := [32]byte{
		// AMSNetId Target
		172, 15, 17, 10, 1, 1,
		// AMSPort Target
		0x21, 0x3, // = 801
		// AMSNetId Target
		172, 15, 17, 15, 1, 1,
		// AMSPort Source
		0xD9, 0x27, // = 10201
		// Command Id,
		0x02, 0x00, // = ADS Read
		// State Flags
		0x05, 0x00, // 	Request
		// Data Length
		0x64, 0x00, 0x00, 0x00, // = 100
		// Error Code
		0x01, 0x00, 0x00, 0x00, // = 1 (ERR_INTERNAL)
		// Invoke Id
		0xC6, 0x22, 0x00, 0x00, // = 8902
	}

	header := parseAMSHeader(b)

	require.NotNil(t, header)

	require.Equal(t, "172.15.17.10.1.1", header.targetNetID.String())
	require.Equal(t, 801, int(header.targetPort))

	require.Equal(t, "172.15.17.15.1.1", header.sourceNetID.String())
	require.Equal(t, 10201, int(header.sourcePort))

	require.Equal(t, 0x0002, int(header.commandID))

	require.Equal(t, 5, int(header.stateFlags))

	require.Equal(t, 100, int(header.dataLength))

	require.Equal(t, ADSErrInternal, int(header.errorCode))

	require.Equal(t, 8902, int(header.invokeID))
}
