package ads

// Symbol .
// See: https://infosys.beckhoff.com/english.php?content=../content/1033/tc3_adsdll2/index.html&id=4279787267115190858
type Symbol struct {
	Name        string `xml:"Name"`
	Comment     string `xml:"Comment"`
	Type        string `xml:"Type"`
	IndexGroup  uint32 `xml:"IGroup"`
	IndexOffset uint32 `xml:"IOffset"`
	Size        uint32 `xml:"BitSize"`
	Flags       uint32 `xml:"Flags"`
}

// See https://infosys.beckhoff.com/english.php?content=../content/1033/tcsample_vc/html/TcAdsDll_API_CPP_Sample10.htm&id=
type adsSymbolInfo struct {
	SymbolCount    uint32
	SymbolLength   uint32
	DataTypeCount  uint32
	DataTypeLength uint32
	ExtraCount     uint32
	ExtraLength    uint32
}
