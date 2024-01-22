package main

import (
	"log"

	"github.com/jobala/tcp/arp"
	"github.com/songgao/packets/ethernet"
	"github.com/songgao/water"
)

func main() {
	config := water.Config{
		DeviceType: water.TAP,
	}
	config.Name = "tap0"

	device, err := water.New(config)
	defer device.Close()

	if err != nil {
		log.Fatal("Failed to create a tap device\n", err)
	}

	var frame ethernet.Frame

	for {
		frame.Resize(1500)
		n, err := device.Read([]byte(frame))
		if err != nil {
			log.Fatal(err)
		}

		frame = frame[:n]

		a := arp.NewArp(device)

		if frame.Ethertype() == ethernet.ARP {
			a.HandleFrame(frame)
		}
	}
}
