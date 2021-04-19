package ads

import "fmt"

// NetID .
type NetID [6]byte

func (n NetID) String() string {
	return fmt.Sprintf("%v.%v.%v.%v.%v.%v", n[0], n[1], n[2], n[3], n[4], n[5])
}

// ParseNetIDFromString .
func ParseNetIDFromString(str string) (NetID, error) {
	var n NetID

	_, err := fmt.Sscanf(str, "%d.%d.%d.%d.%d.%d", &n[0], &n[1], &n[2], &n[3], &n[4], &n[5])
	if err != nil {
		return NetID{}, err
	}

	return n, nil
}

// NetPort .
type NetPort uint16
