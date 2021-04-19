# ADS Go

[![Go Reference](https://pkg.go.dev/badge/github.com/expo21xx/go-ads.svg)](https://pkg.go.dev/github.com/expo21xx/go-ads)

This is a library to interact with Beckhoff ADS/AMS systems.

## Usage

There's an example TwinCAT3 project in `PLCTestProject` and several examples in the `examples` folder.


```go
import (
	ads "github.com/expo21xx/go-ads"
)

client, err := ads.NewClient("192.168.0.5", "192.168.0.5.1.1", ads.AMSPortR0PLCTC3, ads.WithLoadSymbolsOnStart())
if err != nil {
    log.Fatal(err)
}
defer client.Close(context.Background())

err = client.Connect(context.Background())
if err != nil {
    log.Fatal(err)
}

adsState, deviceState, err := client.ReadState(context.Background())
if err != nil {
    log.Fatal(err)
}
fmt.Printf("ADSState: %x | DeviceState: %x\n", adsState, deviceState)
```

## Symbols

The easiest way to work with the data is by the symbol name. These can either be dynamically loaded (and watched for changes) or statically loaded
from a `.tpy` file. These are NOT mutually exclusive, but loading symbols at runtime will override any static symbols with the same name.


Creating a client with `ads.WithLoadSymbolsOnStart()` will load the symbols from the server automatically after connecting.
```go
client, err := ads.NewClient("192.168.0.5", "192.168.0.5.1.1", ads.AMSPortR0PLCTC3, ads.WithLoadSymbolsOnStart())
```

Creating a client with `ads.WithMonitorSymbols()` will create a device notification that will download the new symbols whenever a new version is available.
```go
client, err := ads.NewClient("192.168.0.5", "192.168.0.5.1.1", ads.AMSPortR0PLCTC3, ads.WithMonitorSymbols())
```

Loading symbols from a `.tpy` file can be done using
```go
client, err := ads.NewClient(...)

file, err := os.Open("/path/to/tpy/file")
defer file.Close()
if err != nil {
    log.Fatal(err)
}

err = client.LoadTPYData(file, false) // set to true to load routing information
```
If the second parameter is set to `true` and the `.tpy` file contains routing information, the
client will use the routing information (AMS Net ID and Port) from the file.

Symbols can also be downloaded on demand with the `FetchSymbols` function:

```go
client, err := ads.NewClient(...)
err = client.FetchSymbols(context.Background())
```
