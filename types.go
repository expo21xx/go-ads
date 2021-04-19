package ads

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"math"
	"strings"
	"time"
)

func adsBytesToGoValue(dt string, b []byte) (interface{}, error) {
	switch dt {
	case "BOOL":
		return b[0] == 1, nil
	case "BYTE", "USINT": // USINT = Unsigned Short
		return b[0], nil
	case "SINT":
		return int8(b[0]), nil
	case "UINT", "WORD":
		return binary.LittleEndian.Uint16(b), nil
	case "UDINT", "DWORD":
		return binary.LittleEndian.Uint32(b), nil
	case "INT":
		return int16(binary.LittleEndian.Uint16(b)), nil
	case "DINT":
		return int32(binary.LittleEndian.Uint32(b)), nil
	case "REAL":
		i := binary.LittleEndian.Uint32(b)
		return math.Float32frombits(i), nil
	case "LREAL":
		i := binary.LittleEndian.Uint64(b)
		return math.Float64frombits(i), nil
	case "TIME", "DT":
		i := binary.LittleEndian.Uint32(b)
		return time.Unix(0, int64(uint64(i)*uint64(time.Millisecond))-int64(time.Hour)), nil
	case "TOD": // Time of Day
		i := binary.LittleEndian.Uint32(b)
		return time.Unix(0, int64(uint64(i)*uint64(time.Millisecond))-int64(time.Hour)).Truncate(time.Minute), nil
	}

	if strings.HasPrefix(dt, "WSTRING") {
		l := bytes.Index(b, []byte{0, 0})
		if l < 0 {
			l = len(b)
		}

		r := make([]byte, l/2+1)
		for i, k := 0, 0; i < l; i, k = i+2, k+1 {
			r[k] = b[i]
		}

		return string(r), nil
	}

	if strings.HasPrefix(dt, "STRING") {
		l := bytes.IndexByte(b, 0)
		if l < 0 {
			l = len(b)
		}

		r := make([]byte, l)
		copy(r[:], b[:l])

		return string(r), nil
	}

	return nil, fmt.Errorf("unkown data type %v", dt)
}

func goValueToADSBytes(dt string, val interface{}) ([]byte, error) {
	switch dt {
	case "BOOL":
		b, ok := val.(bool)
		if !ok {
			return nil, fmt.Errorf("provided data %v (%T) is incompatible with data type %v", val, val, dt)
		}

		var r = []byte{0}

		if b {
			r[0] = 1
		}

		return r, nil
	case "BYTE", "USINT": // USINT = Unsigned Short
		b, ok := val.(byte)
		if !ok {
			return nil, fmt.Errorf("provided data %v (%T) is incompatible with data type %v", val, val, dt)
		}
		return []byte{b}, nil
	case "SINT":
		b, ok := val.(int8)
		if !ok {
			return nil, fmt.Errorf("provided data %v (%T) is incompatible with data type %v", val, val, dt)
		}
		return []byte{byte(b)}, nil
	case "UINT", "WORD":
		b, ok := val.(uint16)
		if !ok {
			return nil, fmt.Errorf("provided data %v (%T) is incompatible with data type %v", val, val, dt)
		}
		var r [2]byte
		binary.LittleEndian.PutUint16(r[:], b)
		return r[:], nil
	case "UDINT", "DWORD":
		b, ok := val.(uint32)
		if !ok {
			return nil, fmt.Errorf("provided data %v (%T) is incompatible with data type %v", val, val, dt)
		}
		var r [4]byte
		binary.LittleEndian.PutUint32(r[:], b)
		return r[:], nil
	case "INT":
		b, ok := val.(int16)
		if !ok {
			return nil, fmt.Errorf("provided data %v (%T) is incompatible with data type %v", val, val, dt)
		}
		var r [2]byte
		binary.LittleEndian.PutUint16(r[:], uint16(b))
		return r[:], nil
	case "DINT":
		b, ok := val.(int32)
		if !ok {
			return nil, fmt.Errorf("provided data %v (%T) is incompatible with data type %v", val, val, dt)
		}
		var r [4]byte
		binary.LittleEndian.PutUint32(r[:], uint32(b))
		return r[:], nil
	case "REAL":
		b, ok := val.(float32)
		if !ok {
			return nil, fmt.Errorf("provided data %v (%T) is incompatible with data type %v", val, val, dt)
		}
		var r [4]byte
		binary.LittleEndian.PutUint32(r[:], math.Float32bits(b))
		return r[:], nil
	case "LREAL":
		b, ok := val.(float64)
		if !ok {
			return nil, fmt.Errorf("provided data %v (%T) is incompatible with data type %v", val, val, dt)
		}
		var r [4]byte
		binary.LittleEndian.PutUint64(r[:], math.Float64bits(b))
		return r[:], nil
	}

	if strings.HasPrefix(dt, "STRINGS") || strings.HasPrefix(dt, "WSTRING") {
		b, ok := val.(string)
		if !ok {
			return nil, fmt.Errorf("provided data %v (%T) is incompatible with data type %v", val, val, dt)
		}

		var size int
		_, err := fmt.Scanf("STRINGS(%d)", &size)
		if err != nil {
			return nil, err
		}

		r := make([]byte, size)
		copy(r[:], b)

		return r, nil
	}

	return nil, fmt.Errorf("unkown data type %v", dt)
}
