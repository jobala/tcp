package arp

import (
	"fmt"
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
	mac, _ := net.ParseMAC("8e:8d:d5:66:07:b9")

	reply := &payload{
		hard_type:  Ethernet,
		prot_type:  IPv4,
		hard_size:  MacAddrSize,
		prot_size:  IPv4AddrSize,
		op:         ArpReply,
		sender_mac: mac,
		sender_ip:  []byte{10, 1, 0, 10},
		target_mac: p.sender_mac,
		target_ip:  []byte{10, 22, 14, 78},
	}

	fmt.Printf("%v#", reply.ToEthernetFrame())
	a.Device.Write(reply.ToEthernetFrame())
}
