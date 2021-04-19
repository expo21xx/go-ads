package ads

import (
	"encoding/xml"
	"io"
	"strconv"
)

type tpyProjInfo struct {
	RoutingInfo struct {
		AdsInfo struct {
			NetID string `xml:"NetId"`
			Port  string
		}
	}
	Symbols struct {
		Symbol []Symbol
	}
}

// LoadTPYData .
func (c *Client) LoadTPYData(r io.Reader, loadRoutingInfo bool) error {
	d := xml.NewDecoder(r)
	var tpy tpyProjInfo

	err := d.Decode(&tpy)
	if err != nil {
		return err
	}

	for _, s := range tpy.Symbols.Symbol {
		c.AddSymbol(s)
	}

	if loadRoutingInfo && tpy.RoutingInfo.AdsInfo.NetID != "" {
		targetNetID, err := ParseNetIDFromString(tpy.RoutingInfo.AdsInfo.NetID)
		if err != nil {
			return err
		}
		c.targetNetID = targetNetID

	}

	if loadRoutingInfo && tpy.RoutingInfo.AdsInfo.Port != "" {
		port, err := strconv.ParseUint(tpy.RoutingInfo.AdsInfo.Port, 10, 16)
		if err != nil {
			return err
		}
		c.targetNetPort = NetPort(port)
	}

	return nil
}
