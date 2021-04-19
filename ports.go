package ads

// From: https://github.com/Beckhoff/ADS/blob/master/AdsLib/standalone/AdsDef.h
const (
	ADSTCPServerPort = 0xBF02 // 48898

	AMSPortLogger    = 100
	AMSPortR0RTime   = 200
	AMSPortR0Trace   = (AMSPortR0RTime + 90)
	AMSPortR0IO      = 300
	AMSPortR0SPS     = 400
	AMSPortR0NC      = 500
	AMSPortR0ISG     = 550
	AMSPortR0PCS     = 600
	AMSPortR0PLC     = 801
	AMSPortR0PLCRTS1 = 801
	AMSPortR0PLCRTS2 = 811
	AMSPortR0PLCRTS3 = 821
	AMSPortR0PLCRTS4 = 831
	AMSPortR0PLCTC3  = 851
)
