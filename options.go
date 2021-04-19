package ads

// Option .
type Option interface {
	apply(o *Client)
}

type funcOpt func(c *Client)

func (fo funcOpt) apply(c *Client) {
	fo(c)
}

// WithLoadSymbolsOnStart .
func WithLoadSymbolsOnStart() Option {
	return funcOpt(func(c *Client) {
		c.loadSymbolsOnStart = true
	})
}

// WithMonitorSymbols .
func WithMonitorSymbols() Option {
	return funcOpt(func(c *Client) {
		c.monitorSymbols = true
	})
}

// WithSourceNetID .
func WithSourceNetID(netID NetID) Option {
	return funcOpt(func(c *Client) {
		c.sourceNetID = netID
	})
}

// WithSourceNetPort .
func WithSourceNetPort(port NetPort) Option {
	return funcOpt(func(c *Client) {
		c.sourceNetPort = port
	})
}
