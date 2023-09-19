package helper

import "net"

func IPv4ToUInt32(ipv4Address string) uint32 {
	ip := net.ParseIP(ipv4Address)
	if ip == nil {
		return 0
	}
	ip = ip.To4()
	if ip == nil {
		return 0
	}
	return uint32(ip[0])<<24 + uint32(ip[1])<<16 + uint32(ip[2])<<8 + uint32(ip[3])
}
