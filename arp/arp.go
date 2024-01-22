package arp

import (
	"fmt"

	"os"

	"github.com/songgao/packets/ethernet"
	"github.com/songgao/water"
)

type Arp struct {
	Device *water.Interface
	cache  map[string]string
}

func (a *Arp) HandleFrame(frame ethernet.Frame) {
	p := &payload{}
	p.FromByte(frame.Payload())
	fmt.Printf("%v#", p)

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
		a.reply(*p)
	default:
		fmt.Fprintf(os.Stdout, "unsupported op code: %d\n", []any{p.op}...)
	}
}

func (a *Arp) updateCache(p payload) {
	key := fmt.Sprintf("%d:%d", p.prot_type, p.sender_ip)
	value := fmt.Sprintf("%d:%d:%d", p.prot_type, p.sender_ip, p.sender_mac)
	a.cache[key] = value
}

func (a *Arp) reply(p payload) {
	sender_ip := p.sender_ip
	sender_mac := p.sender_mac

	p.op = ArpReply
	p.sender_mac = p.target_mac
	p.sender_ip = p.target_ip
	p.target_ip = sender_ip
	p.target_mac = sender_mac

	a.Device.Write(p.ToEthernetFrame())
}
