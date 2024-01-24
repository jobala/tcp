package arp

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/songgao/packets/ethernet"
	"github.com/songgao/water"
)

type Arp struct {
	Device *water.Interface
	cache  map[string]string
}

func NewArp(device *water.Interface) *Arp {
	return &Arp{
		Device: device,
		cache:  make(map[string]string, 0),
	}
}

func (a *Arp) HandleFrame(frame ethernet.Frame) {
	p := &payload{}
	p.FromByte(frame.Payload())

    fmt.Println("ARP Request: ", frame)

	if p.hard_type != Ethernet {
		fmt.Printf("unsupported hardware type: %d\n", p.hard_type)
		return
	}
	if p.prot_type != IPv4 {
		fmt.Printf("unsupported protocal type: %d\n", p.prot_type)
	}

	a.updateCache(*p)

	switch p.op {
	case ArpRequest:
		a.reply(p)
	default:
		fmt.Fprintf(os.Stdout, "unsupported op code: %d\n", []any{p.op}...)
	}

}

func (a *Arp) updateCache(p payload) {
	key := fmt.Sprintf("%d:%d", p.prot_type, p.sender_ip)
	value := fmt.Sprintf("%d:%d:%d", p.prot_type, p.sender_ip, p.sender_mac)
	a.cache[key] = value
}

func (a *Arp) reply(p *payload) {
	mac, _ := net.ParseMAC("66:fa:dd:a9:92:48")
    // t_mac, _ := net.ParseMAC("DC:A6:32:8A:80:7B")

	reply := &payload{
		hard_type:  Ethernet,
		prot_type:  IPv4,
		hard_size:  MacAddrSize,
		prot_size:  IPv4AddrSize,
		op:         ArpReply,
		sender_mac: mac,
		sender_ip:  []byte{10, 1, 0, 10},
		target_mac: p.sender_mac,
		target_ip:  p.sender_ip,
	}


    fmt.Println("ARP Reply:  ", reply.ToEthernetFrame())
	if _, err := a.Device.Write(reply.ToEthernetFrame()); err != nil {
		log.Fatal("something went wrong", err)
	}

	log.Println("frame sent successfully")
}
