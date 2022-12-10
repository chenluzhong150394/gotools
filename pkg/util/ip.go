package util

import (
	"errors"
	"net"
)

// ExternalIP 获取ip
func ExternalIP() (ip net.IP, err error) {
	var address []net.Addr
	netInterface, err := net.Interfaces()
	if err != nil {
		return
	}

	for _, n := range netInterface {
		if n.Flags&net.FlagUp == 0 {
			continue // interface down
		}
		if n.Flags&net.FlagLoopback != 0 {
			continue // logback interface
		}
		address, err = n.Addrs()
		if err != nil {
			return
		}
		for _, addr := range address {
			ip = getIpFromAddr(addr)
			if ip == nil {
				continue
			}
			return
		}
	}
	err = errors.New("connected to the network")
	return
}

//获取ip
func getIpFromAddr(addr net.Addr) net.IP {
	var ip net.IP
	switch v := addr.(type) {
	case *net.IPNet:
		ip = v.IP
	case *net.IPAddr:
		ip = v.IP
	}
	if ip == nil || ip.IsLoopback() {
		return nil
	}
	ip = ip.To4()
	if ip == nil {
		return nil // not an ipv4 address
	}

	return ip
}

// GetHost 获取当前本机的ip
func GetHost() (host string) {
	if ip, err := ExternalIP(); err == nil {
		host = ip.String()
	} else {
		host = "localhost"
	}
	return
}
