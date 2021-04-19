package ads

// from  https://golang.org/src/syscall/types_windows.go
type filetime struct {
	lowDateTime  uint32
	highDateTime uint32
}

// Nanoseconds returns Filetime ft in nanoseconds
// since Epoch (00:00:00 UTC, January 1, 1970).
func (ft *filetime) nanoseconds() int64 {
	// 100-nanosecond intervals since January 1, 1601
	nsec := int64(ft.highDateTime)<<32 + int64(ft.lowDateTime)
	// change starting time to the Epoch (00:00:00 UTC, January 1, 1970)
	nsec -= 116444736000000000
	// convert into nanoseconds
	nsec *= 100
	return nsec
}

func nsecToFiletime(nsec int64) (ft filetime) {
	// convert into 100-nanosecond
	nsec /= 100
	// change starting time to January 1, 1601
	nsec += 116444736000000000
	// split into high / low
	ft.lowDateTime = uint32(nsec & 0xffffffff)
	ft.highDateTime = uint32(nsec >> 32 & 0xffffffff)
	return ft
}
