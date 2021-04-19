package ads

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestReadDeviceInfoCmd(t *testing.T) {
	var _ = (Cmd)(&ReadDeviceInfoCmdRequest{})

	expected := &ReadDeviceInfoCmdResponse{
		Result:       ADSErrNoError,
		MajorVersion: 3,
		MinorVersion: 2,
		VersionBuild: 20,
		DeviceName:   "Test",
	}

	actual := &ReadDeviceInfoCmdResponse{}
	err := actual.fromBytes([]byte{
		// Error Code
		0x00, 0x00, 0x00, 0x00, // = 1 (ERR_NOERROR)
		// MajorVersion
		0x03,
		// MinorVersion
		0x02,
		// VersionBuild
		0x14, 0x0,
		84, 101, 115, 116, // = "Test"
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
	})
	require.Nil(t, err)

	require.Equal(t, expected, actual)
}

func TestReadCmd(t *testing.T) {
	var _ = (Cmd)(&ReadCmdRequest{})

	req := &ReadCmdRequest{
		IndexGroup:  0x4,
		IndexOffset: 0x1,
		Length:      10,
	}

	actualReq := &ReadCmdRequest{}
	err := actualReq.FromBytes(req.Bytes())
	require.Nil(t, err)
	require.Equal(t, req, actualReq)

	resExpected := &ReadCmdResponse{
		Result: ADSErrNOIO,
		Data:   []byte("Test"),
	}

	resActual := &ReadCmdResponse{}
	err = resActual.fromBytes([]byte{
		// Error Code
		0x0A, 0x00, 0x00, 0x00, // = ERR_NOIO
		0x04, 0, 0, 0, // = 4
		84, 101, 115, 116, // = "Test"
	})
	require.Nil(t, err)

	require.Equal(t, resExpected, resActual)
}

func TestWriteCmd(t *testing.T) {
	var _ = (Cmd)(&WriteCmdRequest{})

	req := &WriteCmdRequest{
		IndexGroup:  0x4,
		IndexOffset: 0x1,
		Data:        []byte{84, 101, 115, 116}, // = "Test",
	}

	actualReq := &WriteCmdRequest{}
	err := actualReq.FromBytes(req.Bytes())
	require.Nil(t, err)
	require.Equal(t, req, actualReq)

	resExpected := &WriteCmdResponse{
		Result: ADSErrNOIO,
	}

	resActual := &WriteCmdResponse{}
	err = resActual.fromBytes([]byte{
		0x0A, 0x00, 0x00, 0x00, // = ERR_NOIO
	})
	require.Nil(t, err)

	require.Equal(t, resExpected, resActual)
}

func TestReadStateCmd(t *testing.T) {
	var _ = (Cmd)(&ReadStateCmdRequest{})

	resExpected := &ReadStateCmdResponse{
		Result:      ADSErrNOIO,
		ADSState:    ADSStateRun,
		DeviceState: 2,
	}

	resActual := &ReadStateCmdResponse{}
	err := resActual.fromBytes([]byte{
		0x0A, 0x00, 0x00, 0x00, // = ERR_NOIO
		0x05, 0x0, // = ADSStateRun
		// DeviceState
		0x02, 0x0, // = 2
	})
	require.Nil(t, err)

	require.Equal(t, resExpected, resActual)
}

func TestWriteControlCmd(t *testing.T) {
	var _ = (Cmd)(&WriteControlCmdRequest{})

	req := &WriteControlCmdRequest{
		ADSState:    ADSStateRun,
		DeviceState: 2,
		Data:        []byte{84, 101, 115, 116}, // = "Test",
	}

	actualReq := &WriteControlCmdRequest{}
	err := actualReq.FromBytes(req.Bytes())
	require.Nil(t, err)
	require.Equal(t, req, actualReq)

	resExpected := &WriteControlCmdResponse{
		Result: ADSErrNOIO,
	}

	resActual := &WriteControlCmdResponse{}
	err = resActual.fromBytes([]byte{
		0x0A, 0x00, 0x00, 0x00, // = ERR_NOIO
	})
	require.Nil(t, err)

	require.Equal(t, resExpected, resActual)
}

func TestAddDeviceNotificationCmd(t *testing.T) {
	var _ = (Cmd)(&AddDeviceNotificationCmdRequest{})

	req := &AddDeviceNotificationCmdRequest{
		IndexGroup:       0x4,
		IndexOffset:      0x1,
		Length:           4,
		TransmissionMode: ADSTransServerCycle,
		MaxDelay:         10,
		CycleTime:        100,
	}

	actualReq := &AddDeviceNotificationCmdRequest{}
	err := actualReq.FromBytes(req.Bytes())
	require.Nil(t, err)
	require.Equal(t, req, actualReq)

	resExpected := &AddDeviceNotificationCmdResponse{
		Result: ADSErrNOIO,
	}

	resActual := &AddDeviceNotificationCmdResponse{}
	err = resActual.fromBytes([]byte{
		0x0A, 0x00, 0x00, 0x00, // = ERR_NOIO
		0x00, 0x00, 0x00, 0x00,
	})
	require.Nil(t, err)

	require.Equal(t, resExpected, resActual)
}

func TestDeleteDeviceNotificationCmd(t *testing.T) {
	var _ = (Cmd)(&DeleteDeviceNotificationCmdRequest{})

	req := &DeleteDeviceNotificationCmdRequest{
		NotificationHandle: 0x4,
	}

	actualReq := &DeleteDeviceNotificationCmdRequest{}
	err := actualReq.FromBytes(req.Bytes())
	require.Nil(t, err)
	require.Equal(t, req, actualReq)

	resExpected := &DeleteDeviceNotificationCmdResponse{
		Result: ADSErrNOIO,
	}

	resActual := &DeleteDeviceNotificationCmdResponse{}
	err = resActual.fromBytes([]byte{
		0x0A, 0x00, 0x00, 0x00,
	})
	require.Nil(t, err)

	require.Equal(t, resExpected, resActual)
}

func TestReadWriteCmd(t *testing.T) {
	var _ = (Cmd)(&ReadWriteCmdRequest{})

	req := &ReadWriteCmdRequest{
		IndexGroup:  0x4,
		IndexOffset: 0x1,
		Data:        []byte{84, 101, 115, 116}, // = "Test",
	}

	actualReq := &ReadWriteCmdRequest{}
	err := actualReq.FromBytes(req.Bytes())
	require.Nil(t, err)
	require.Equal(t, req, actualReq)

	resExpected := &ReadWriteCmdResponse{
		Result: ADSErrNOIO,
		Data:   []byte{84, 101, 115, 116}, // = "Test",
	}

	resActual := &ReadWriteCmdResponse{}
	err = resActual.fromBytes([]byte{
		0x0A, 0x00, 0x00, 0x00, // = ERR_NOIO
		0x04, 0, 0, 0,
		84, 101, 115, 116, // = "Test"
	})
	require.Nil(t, err)

	require.Equal(t, resExpected, resActual)
}

func TestDeviceNotificationCmd(t *testing.T) {
	var _ = (Cmd)(&DeviceNotificationCmdRequest{})

	ts := time.Date(2021, 01, 10, 0, 0, 0, 0, time.UTC)

	req := &DeviceNotificationCmdRequest{
		Stamps: []StampHeader{
			{
				Timestamp: ts,
				Samples: []Sample{
					{
						NotificationHandle: 0xA4,
						Data:               []byte{69, 110, 116, 114, 121, 49}, // = "Entry1"
					},
					{
						NotificationHandle: 0xA4,
						Data:               []byte{69, 110, 116, 114, 121, 50}, // = "Entry2"
					},
				},
			},
			{
				Timestamp: ts,
				Samples: []Sample{
					{
						NotificationHandle: 0xA4,
						Data:               []byte{69, 110, 116, 114, 121, 51}, // = "Entry3"
					},
				},
			},
			{
				Timestamp: ts,
				Samples: []Sample{
					{
						NotificationHandle: 0xA4,
						Data:               []byte{69, 110, 116, 114, 121, 52}, // = "Entry4"
					},
					{
						NotificationHandle: 0xA4,
						Data:               []byte{69, 110, 116, 114, 121, 53}, // = "Entry5"
					},
				},
			},
		},
	}

	actualReq := &DeviceNotificationCmdRequest{}
	err := actualReq.FromBytes(req.Bytes())

	require.Nil(t, err)
	require.Equal(t, req, actualReq)
}
