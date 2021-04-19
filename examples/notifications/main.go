package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	ads "github.com/expo21xx/go-ads"
)

func main() {
	ip := flag.String("ip", "192.168.0.5", "ip address of the AMS router")
	netid := flag.String("netid", "192.168.0.5.1.1", "target AMS NetID")
	port := flag.Int("port", ads.AMSPortR0PLCTC3, "target AMS port")
	timeout := flag.Int("timeout", 15, "timeout in s")
	names := flag.String("names", "", "comma separated list of names to listen for changes for")

	flag.Parse()

	client, err := ads.NewClient(*ip, *netid, *port, ads.WithLoadSymbolsOnStart(), ads.WithMonitorSymbols())
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close(context.Background())

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(*timeout))
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	adsState, deviceState, err := client.ReadState(context.Background())
	if err != nil {
		client.Close(context.Background())
		log.Fatal(err)
	}
	fmt.Printf("ADSState: %x | DeviceState: %x\n", adsState, deviceState)

	for _, name := range strings.Split(*names, ",") {
		stop, err := client.AddDeviceNotificationHandlerByName(context.Background(), name, nil, func(val interface{}) {
			fmt.Printf("New value for %v: %v\n", name, val)
		})
		defer stop()
		if err != nil {
			log.Println(err)
		}
	}

	shutdownHandler(func() error {
		return client.Close(context.Background())
	})
}

func shutdownHandler(handler func() error) {
	// buffered channel because the signal module requires it
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)

	select {
	case sig := <-shutdown:
		log.Printf("received signal %v: triggering shutdown\n", sig)

		err := handler()

		// wait for current requests to finish
		switch {
		case sig == syscall.SIGSTOP:
			log.Fatal("SIGSTOP caused shutdown")
		case err != nil:
			log.Fatal(err)
		}
	}
}
