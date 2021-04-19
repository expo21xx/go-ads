package ads

// From https://github.com/Beckhoff/ADS/blob/master/AdsLib/standalone/AdsDef.h
const (
	ADSIndexGroupSymtab  uint32 = 0xF000
	ADSIndexGroupSymname uint32 = 0xF001
	ADSIndexGroupSymval  uint32 = 0xF002

	ADSIndexGroupSymHndByName    uint32 = 0xF003
	ADSIndexGroupSymValByName    uint32 = 0xF004
	ADSIndexGroupSymValByHnd     uint32 = 0xF005
	ADSIndexGroupSymReleaseHnd   uint32 = 0xF006
	ADSIndexGroupSymInfoByName   uint32 = 0xF007
	ADSIndexGroupSymVersion      uint32 = 0xF008
	ADSIndexGroupSymInfoByNameEx uint32 = 0xF009

	ADSIndexGroupSymDownload    uint32 = 0xF00A
	ADSIndexGroupSymUpload      uint32 = 0xF00B
	ADSIndexGroupSymUploadInfo  uint32 = 0xF00C
	ADSIndexGroupSymDownload2   uint32 = 0xF00D
	ADSIndexGroupSymDTUpload    uint32 = 0xF00E
	ADSIndexGroupSymUploadInfo2 uint32 = 0xF00F

	// notification of named handle
	ADSIndexGroupSymNote uint32 = 0xF010
)
