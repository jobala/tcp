package main

import (
	"log"

	"github.com/songgao/packets/ethernet"
	"github.com/songgao/water"
)

func main() {
	config := water.Config{
		DeviceType: water.TAP,
	}
	config.Name = "tcp"

	device, err := water.New(config)
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
		log.Printf("Dst: %s\n", frame.Destination())
		log.Printf("Src: %s\n", frame.Source())
		log.Printf("Ethertype: % x\n", frame.Ethertype())
		log.Printf("Payload: % x\n", frame.Payload())
	}
}
