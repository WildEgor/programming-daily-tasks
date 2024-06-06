package cidr_firewall

import (
	"errors"
	"net"
	"strconv"
	"strings"
)

const (
	ALLOWED = "ALLOWED"
	DENY    = "DENY"
)

func IsAllowed(rules [][2]string, ip string) bool {
	ipAddr := net.ParseIP(ip)
	if ipAddr == nil {
		// Invalid IP address
		return false
	}

	for _, rule := range rules {
		cidr := rule[0]
		action := rule[1]

		// Handle single IP address
		if !strings.Contains(cidr, "/") {
			if ipAddr.Equal(net.ParseIP(cidr)) {
				return action == ALLOWED
			}
			continue
		}

		// Handle CIDR notation
		_, ipNet, err := net.ParseCIDR(cidr)
		if err != nil {
			// Invalid CIDR notation
			continue
		}

		if ipNet.Contains(ipAddr) {
			return action == ALLOWED
		}
	}

	// Default action if no rules match
	return false
}

func strIpToBinary(ip string) (uint32, error) {
	parts := strings.Split(ip, ".") // 192.168.1.1 ->
	if len(parts) != 4 {
		return 0, errors.New("incorrect ip")
	}

	var result uint32
	for i := 0; i < 4; i++ {
		part, err := strconv.Atoi(parts[i])
		if err != nil {
			return 0, errors.New("incorrect ip")
		}

		// This part convert ip each dec part representation to binary
		// 1) Shift to left
		// 2) For example, if result was 00000001 (in binary), shifting it 8 positions to the left would result in 0000000100000000.
		// 3) For example, if result is 0000000100000000 (after shifting) and part is 00000011 (in binary), the OR operation would result in 0000000100000011.
		result = result<<8 | uint32(part)
	}

	return result, nil
}
