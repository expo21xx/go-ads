package ads

import "fmt"

// From https://github.com/Beckhoff/ADS/blob/master/AdsLib/standalone/AdsDef.h
const (
	// ADSErrNoError (ERR_NOERROR): No error.
	ADSErrNoError = 0x00

	// ADSErrInternal (ERR_INTERNAL): Internal error.
	ADSErrInternal = 0x1

	// ADSErrTargetPortNotFound (ERR_TARGETPORTNOTFOUND): Target port not found – ADS server is not started or is not reachable.
	ADSErrTargetPortNotFound = 0x6

	// Invalid AMS port.
	ADSErrInvalidAMSPort = 0x18

	// ADSErrNOIO (ERR_NOIO): No IO.
	ADSErrNOIO = 0xA

	errADSErrs = 0x0700

	// Error class < device error >
	ADSErrDeviceError = (0x00 + errADSErrs)

	// Service is not supported by server
	ADSErrDeviceSrvNotSupp = (0x01 + errADSErrs)

	// invalid indexGroup
	ADSErrDeviceInvalidGrp = (0x02 + errADSErrs)

	// invalid indexOffset
	ADSErrDeviceInvaliIoffset = (0x03 + errADSErrs)

	// reading/writing not permitted
	ADSErrDeviceInvalidAccess = (0x04 + errADSErrs)

	// parameter size not correct
	ADSErrDeviceInvalidSize = (0x05 + errADSErrs)

	// invalid parameter value(s)
	ADSErrDeviceInvalidData = (0x06 + errADSErrs)

	// device is not in a ready state
	ADSErrDeviceNotReady = (0x07 + errADSErrs)

	// device is busy
	ADSErrDeviceBusy = (0x08 + errADSErrs)

	// invalid context (must be InWindows)
	ADSErrDeviceInvalidcontext = (0x09 + errADSErrs)

	// invalid parameter value(s)ut of memory
	ADSErrDeviceNoMemory = (0x0A + errADSErrs)

	// invalid parameter value(s)
	ADSErrDeviceInvalidParm = (0x0B + errADSErrs)

	// not found (files, ...)
	ADSErrDeviceNotFound = (0x0C + errADSErrs)

	// syntax error in comand or file
	ADSErrDeviceSyntax = (0x0D + errADSErrs)

	// objects do not match
	ADSErrDeviceIncompatible = (0x0E + errADSErrs)

	// object already exists
	ADSErrDeviceExists = (0x0F + errADSErrs)

	// symbol not found
	ADSErrDeviceSymbolNotFound = (0x10 + errADSErrs)

	// symbol version invalid, possibly caused by an 'onlinechange' -> try to release handle and get a new one
	ADSErrDeviceSymbolVersionInvalid = (0x11 + errADSErrs)

	// server is in invalid state*/
	ADSErrDeviceInvalidState = (0x12 + errADSErrs)

	// AdsTransMode not supported
	ADSErrDeviceTransmodeNotSupp = (0x13 + errADSErrs)

	// Notification handle is invalid, possibly caussed by an 'onlinechange' -> try to release handle and get a new one
	ADSErrDeviceNotifyHndInvalid = (0x14 + errADSErrs)

	// Notification client not registered*/
	ADSErrDeviceClientUnknown = (0x15 + errADSErrs)

	// no more notification handles
	ADSErrDeviceNomoreHdls = (0x16 + errADSErrs)

	// size for watch to big
	ADSErrDeviceInvalidWatchSize = (0x17 + errADSErrs)

	// device not initialized
	ADSErrDeviceNotInit = (0x18 + errADSErrs)

	// device has a timeout
	ADSErrDeviceTimeout = (0x19 + errADSErrs)

	// query interface failed
	ADSErrDeviceNoInterface = (0x1A + errADSErrs)

	// wrong interface required
	ADSErrDeviceInvalidInterface = (0x1B + errADSErrs)

	// class ID is invalid
	ADSErrDeviceInvalidClsID = (0x1C + errADSErrs)

	// object ID is invalid
	ADSErrDeviceInvalidObjID = (0x1D + errADSErrs)

	// request is pending
	ADSErrDevicePending = (0x1E + errADSErrs)

	// request is aborted
	ADSErrDeviceAborted = (0x1F + errADSErrs)

	// signal warningequest is aborted
	ADSErrDeviceWarning = (0x20 + errADSErrs)

	// invalid array index
	ADSErrDeviceInvalidArrayIdx = (0x21 + errADSErrs)

	// symbol not active, possibly caussed by an 'onlinechange' -> try to release handle and get a new one
	ADSErrDeviceSymbolNotActive = (0x22 + errADSErrs)

	// access denied
	ADSErrDeviceAccessDenied = (0x23 + errADSErrs)

	// no license found -> Activate license for TwinCAT 3 function
	ADSErrDeviceLicenseNotFound = (0x24 + errADSErrs)

	// license expired
	ADSErrDeviceLicenseExpired = (0x25 + errADSErrs)

	// license exceeded
	ADSErrDeviceLicenseExceeded = (0x26 + errADSErrs)

	// license invalid
	ADSErrDeviceLicenseInvalid = (0x27 + errADSErrs)

	// license invalid system id
	ADSErrDeviceLicenseSystemID = (0x28 + errADSErrs)

	// license not time limited
	ADSErrDeviceLicenseNoTimeLimit = (0x29 + errADSErrs)

	// license issue time in the future
	ADSErrDeviceLicenseFutureIssue = (0x2A + errADSErrs)

	// license time period to long
	ADSErrDeviceLicenseTimeToLong = (0x2B + errADSErrs)

	// exception in device specific code -> Check each device transistions
	ADSErrDeviceException = (0x2C + errADSErrs)

	// license file read twice
	ADSErrDeviceLicenseDuplicated = (0x2D + errADSErrs)

	// invalid signature
	ADSErrDeviceSignatureInvalid = (0x2E + errADSErrs)

	// public key certificate
	ADSErrDeviceCertificateInvalid = (0x2F + errADSErrs)

	// Error class < client error >
	ADSErrClientError = (0x40 + errADSErrs)

	// invalid parameter at service call
	ADSErrClientInvalidParm = (0x41 + errADSErrs)

	// callling list is empty
	ADSErrClientListEmpty = (0x42 + errADSErrs)

	// var connection already in use
	ADSErrClientVarUsed = (0x43 + errADSErrs)

	// invoke id in use
	ADSErrClientduplInvokeID = (0x44 + errADSErrs)

	// timeout elapsed -> Check ADS routes of sender and receiver and your [firewall setting](http://infosys.beckhoff.com/content/1033/tcremoteaccess/html/tcremoteaccess_firewall.html?id=12027)
	ADSErrClientSyncTimeout = (0x45 + errADSErrs)

	// error in win32 subsystem
	ADSErrClientW32Error = (0x46 + errADSErrs)

	// invalid client timeout value
	ADSErrClientTimeoutInvalid = (0x47 + errADSErrs)

	// ads dll
	ADSErrClientportNotOpen = (0x48 + errADSErrs)

	// ads dll
	ADSErrClientNoAMSAddr = (0x49 + errADSErrs)

	// Internal error in ads sync
	ADSErrClientSyncInternal = (0x50 + errADSErrs)

	// hash table overflow
	ADSErrClientAddHash = (0x51 + errADSErrs)

	// key not found in hash table
	ADSErrClientRemoveHash = (0x52 + errADSErrs)

	// no more symbols in cache
	ADSErrClientNoMoreSym = (0x53 + errADSErrs)

	// invalid response received
	ADSErrClientSyncResInvalid = (0x54 + errADSErrs)

	// sync port is locked
	ADSErrClientSyncPortLocked = (0x55 + errADSErrs)
)

// ErrCodeToString .
func ErrCodeToString(c uint32) string {
	switch c {
	case ADSErrNoError:
		return "no error"
	case ADSErrInternal:
		return "internal error"
	case ADSErrTargetPortNotFound:
		return "target port not found: ADS server is not started or is not reachable"
	case ADSErrInvalidAMSPort:
		return "invalid AMS port"
	case ADSErrDeviceSrvNotSupp:
		return "service is not supported by the server"
	case ADSErrDeviceInvalidGrp:
		return "invalid index group"
	case ADSErrDeviceInvalidSize:
		return "parameter size not correct"
	case ADSErrDeviceSymbolNotFound:
		return "symbol not found"
	case ADSErrDeviceInvalidWatchSize:
		return "notification size too large"
	}

	return fmt.Sprintf("unkown error code %v", c)
}
