package mock

import (
	"crypto/rand"
	"net"
)

// RandIPv4 Returns a random IPv4 address
func RandIPv4() string {
	newip := make(net.IP, len(net.IPv4zero))
	copy(newip, net.IPv4zero)
	rand.Read(newip[len(newip)-4:])
	return newip.String()
}

// RandIPv6 Returns a random IPv6 address
func RandIPv6() string {
	newip := make(net.IP, len(net.IPv6zero))
	rand.Read(newip[:])
	return newip.String()
}
