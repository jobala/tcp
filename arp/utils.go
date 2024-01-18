package arp

import "time"

const (
	macAddrSize       = 6
	ethARP            = 0x806
	protIPv4          = 0x800
	arpRequest  uint8 = 1
	arpReply    uint8 = 2
)

const requestTimeout = 500 * time.Millisecond
