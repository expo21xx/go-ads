package ads

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParseAMSTCPHeader(t *testing.T) {
	b := [6]byte{0, 0, 40, 0, 0, 0} // taken from a real request

	header := parseAMSTCPHeader(b)

	require.Equal(t, 40, int(header.length))
}
