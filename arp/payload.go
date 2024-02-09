package arp

import (
	"encoding/binary"
)

type payload struct {
	hard_type  uint16
	prot_type  uint16
	hard_size  uint8
	prot_size  uint8
	op         uint16
	sender_mac []byte
	sender_ip  []byte
	target_mac []byte
	target_ip  []byte
}

func NewPayload(arpOp uint16) *payload {
	return &payload{
		hard_type: Ethernet,
		prot_type: IPv4,
		hard_size: MacAddrSize,
		prot_size: IPv4AddrSize,
		op:        arpOp,
	}
}

func (p *payload) ToByte() []byte {
	packet := make([]byte, 28)

	binary.BigEndian.PutUint16(packet[0:2], p.hard_type)
	binary.BigEndian.PutUint16(packet[2:4], p.prot_type)
	packet[4] = p.hard_size
	packet[5] = p.prot_size
	binary.BigEndian.PutUint16(packet[6:8], p.op)
	copy(packet[8:14], p.sender_mac)
	copy(packet[14:18], p.sender_ip)
	copy(packet[18:24], p.target_mac)
	copy(packet[24:28], p.target_ip)

	return packet
}

func (p *payload) FromByte(data []byte) *payload {
	p.hard_type = uint16(uint16(data[0])<<8 | uint16(data[1]))
	p.prot_type = uint16(uint16(data[2])<<8 | uint16(data[3]))
	p.hard_size = uint8(data[4])
	p.prot_size = uint8(data[5])
	p.op = uint16(uint16(data[6])<<8 | uint16(data[7]))
	p.sender_mac = data[8:14]
	p.sender_ip = data[14:18]
	p.target_mac = data[18:24]
	p.target_ip = data[24:28]
	return p
}

func (p *payload) ToEthernetFrame() []byte {
	frame := make([]byte, 64)
	etherType := 0x0806 // ARP, you can replace this with other EtherType values

	copy(frame[0:6], p.target_mac)
	copy(frame[6:12], p.sender_mac)
	binary.BigEndian.PutUint16(frame[12:14], uint16(etherType))
	copy(frame[14:43], p.ToByte())

	return frame
}
