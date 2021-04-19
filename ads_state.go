package ads

// From the C++ API
// https://infosys.beckhoff.com/english.php?content=../content/1033/tc3_adsdll2/index.html&id=4279787267115190858
const (
	ADSStateInvalid      = 0
	ADSStateidle         = 1
	ADSStatereset        = 2
	ADSStateInit         = 3
	ADSStateStart        = 4
	ADSStateRun          = 5
	ADSStateStop         = 6
	ADSStateSavecfg      = 7
	ADSStateLoadcfg      = 8
	ADSStatePowerfailure = 9
	ADSStatePowergood    = 10
	ADSStateError        = 11
	ADSStateShutdown     = 12
	ADSStateSuspend      = 13
	ADSStateResume       = 14
	// system is in config mode
	ADSStateConfig = 15
	// system should restart in config mode
	ADSStateReconfig  = 16
	ADSStateMaxstates = 17
)
