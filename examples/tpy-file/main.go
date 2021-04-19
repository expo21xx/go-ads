package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	ads "github.com/expo21xx/go-ads"
)

func main() {
	ip := flag.String("ip", "192.168.0.5", "ip address of the AMS router")
	timeout := flag.Int("timeout", 5, "timeout in s")
	filename := flag.String("file", "PLCTestProject/PLCTestProject/PLC/PLC.tpy", ".tpy file to read symbol data from")
	readValue := flag.String("read", "", "value to read")

	flag.Parse()

	client, err := ads.NewClient(*ip, "", 0)
	if err != nil {
		log.Fatal(err)
	}

	file, err := os.Open(*filename)
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}

	err = client.LoadTPYData(file, true)
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
		log.Fatal(err)
	}
	fmt.Printf("ADSState: %x | DeviceState: %x\n", adsState, deviceState)

	if *readValue != "" {
		value, err := client.ReadByName(context.Background(), *readValue)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Value for %v: %v\n", *readValue, value)
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
