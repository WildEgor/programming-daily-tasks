package cidr_firewall

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_IsAllowed(t *testing.T) {
	cidr := [][2]string{
		{"192.168.1.0/24", ALLOWED},
		{"10.0.0.0./16", DENY},
		{"8.8.8.8", ALLOWED},
	}
	ip := "192.168.1.10"
	assert.Equal(t, IsAllowed(cidr, ip), true)
}

func Test_strIpToBinary(t *testing.T) {
	in := "192.168.1.1"
	result, _ := strIpToBinary(in)
	assert.Equal(t, result, uint32(0xc0a80101)) // 11000000101010000000000100000001
}
