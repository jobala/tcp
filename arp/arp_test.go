package arp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFromByte(t *testing.T) {
	p := &payload{}
	p.FromByte([]byte{
		0, 1,
		0x08, 0x00,
		6,
		4,
		0, 1,
		1, 2, 3, 4, 5, 6,
		192, 168, 1, 1,
		0, 0, 0, 0, 0, 0,
		192, 168, 1, 2,
	})

	assert.Equal(t, p, &payload{
		hard_type:  1,
		prot_type:  0x0800,
		hard_size:  6,
		prot_size:  4,
		op:         1,
		sender_mac: []byte{1, 2, 3, 4, 5, 6},
		sender_ip:  []byte{192, 168, 1, 1},
		target_mac: []byte{0, 0, 0, 0, 0, 0},
		target_ip:  []byte{192, 168, 1, 2},
	})
}

func TestToByte(t *testing.T) {
	p := &payload{
		hard_type:  1,
		prot_type:  0x0800,
		hard_size:  6,
		prot_size:  4,
		op:         1,
		sender_mac: []byte{1, 2, 3, 4, 5, 6},
		sender_ip:  []byte{192, 168, 1, 1},
		target_mac: []byte{0, 0, 0, 0, 0, 0},
		target_ip:  []byte{192, 168, 1, 2},
	}

	assert.Equal(t, p.ToByte(), []byte{
		0, 1,
		0x08, 0x00,
		6,
		4,
		0, 1,
		1, 2, 3, 4, 5, 6,
		192, 168, 1, 1,
		0, 0, 0, 0, 0, 0,
		192, 168, 1, 2,
	})
}
