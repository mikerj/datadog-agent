package enrichment

import (
	"net"
	"strconv"
)

func FormatMask(ipAddr []byte, maskRawValue uint32) string {
	maskSuffix := "/" + strconv.Itoa(int(maskRawValue))

	ip := net.IP(ipAddr)
	if ip == nil {
		return maskSuffix
	}

	var maskBitsLen int
	// Using ip.To4() to test for ipv4
	// More info: https://stackoverflow.com/questions/40189084/what-is-ipv6-for-localhost-and-0-0-0-0
	if ip.To4() != nil {
		maskBitsLen = 32
	} else {
		maskBitsLen = 128
	}

	mask := net.CIDRMask(int(maskRawValue), maskBitsLen)
	if mask == nil {
		return maskSuffix
	}
	maskedIP := ip.Mask(mask)
	if maskedIP == nil {
		return maskSuffix
	}
	return maskedIP.String() + maskSuffix
}
