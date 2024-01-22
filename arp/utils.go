package arp

import "time"

const (
	Ethernet            = 1
	MacAddrSize         = 6
	IPv4AddrSize        = 4
	IPv4                = 0x0800
	ArpRequest   uint16 = 1
	ArpReply     uint16 = 2
)

const RequestTimeout = 500 * time.Millisecond
const IP = "10.1.0.10"
