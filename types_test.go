package ads

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestADSBytesToGoValue(t *testing.T) {
	v, err := adsBytesToGoValue("BOOL", []byte{1})
	require.Nil(t, err)
	require.Equal(t, true, v)

	v, err = adsBytesToGoValue("BYTE", []byte{9})
	require.Nil(t, err)
	require.Equal(t, uint8(9), v)

	v, err = adsBytesToGoValue("USINT", []byte{12})
	require.Nil(t, err)
	require.Equal(t, uint8(12), v)

	v, err = adsBytesToGoValue("SINT", []byte{235})
	require.Nil(t, err)
	require.Equal(t, int8(-21), v)

	v, err = adsBytesToGoValue("UINT", []byte{32, 0})
	require.Nil(t, err)
	require.Equal(t, uint16(32), v)

	v, err = adsBytesToGoValue("WORD", []byte{54, 0})
	require.Nil(t, err)
	require.Equal(t, uint16(54), v)

	v, err = adsBytesToGoValue("UDINT", []byte{44, 0, 0, 0})
	require.Nil(t, err)
	require.Equal(t, uint32(44), v)

	v, err = adsBytesToGoValue("DWORD", []byte{89, 0, 0, 0})
	require.Nil(t, err)
	require.Equal(t, uint32(89), v)

	v, err = adsBytesToGoValue("INT", []byte{233, 255})
	require.Nil(t, err)
	require.Equal(t, int16(-23), v)

	v, err = adsBytesToGoValue("DINT", []byte{162, 162, 255, 255})
	require.Nil(t, err)
	require.Equal(t, int32(-23902), v)

	v, err = adsBytesToGoValue("REAL", []byte{195, 240, 1, 64})
	require.Nil(t, err)
	require.Equal(t, float32(2.03032), v)

	v, err = adsBytesToGoValue("LREAL", []byte{187, 196, 198, 154, 197, 207, 16, 192})
	require.Nil(t, err)
	require.Equal(t, float64(-4.20290223921), v)

	v, err = adsBytesToGoValue("STRING(80)", []byte{83, 116, 114, 105, 110, 103, 32, 116, 101, 115, 116, 32, 118, 97, 114, 32, 118, 97, 108, 117, 101, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0})
	require.Nil(t, err)
	require.Equal(t, "String test var value", v)

	v, err = adsBytesToGoValue("WSTRING(80)", []byte{87, 0, 83, 0, 116, 0, 114, 0, 105, 0, 110, 0, 103, 0, 32, 0, 116, 0, 101, 0, 115, 0, 116, 0, 32, 0, 118, 0, 97, 0, 114, 0, 32, 0, 118, 0, 97, 0, 108, 0, 117, 0, 101, 0, 0, 0, 131, 180, 221, 6, 0, 91, 94, 95, 89, 90, 88, 72, 131, 196, 16, 201, 195, 102, 144, 144, 6, 0, 51, 0, 46, 9, 86, 9, 72, 12, 248, 12, 109, 13, 142, 15, 211, 16, 86, 17, 181, 17, 9, 18, 112, 18, 190, 18, 39, 19, 133, 19, 46, 21, 200, 21, 201, 22, 142, 26, 197, 27, 143, 28, 89, 29, 62, 115, 94, 115, 130, 115, 82, 184, 130, 184, 178, 184, 226, 184, 18, 185, 66, 185, 114, 185, 162, 185, 210, 185, 2, 186, 50, 186, 98, 186, 2, 187, 106, 187, 154, 187, 202, 187, 250, 187, 42, 188, 90, 188, 138, 188, 186, 188, 0, 0})
	require.Nil(t, err)
	require.Equal(t, "WString test var value", v)
}
