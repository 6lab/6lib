package labbin

import (
	"net"
	"os"
	"strings"
)

// Info contains information about a binary
type Info struct {
	Path     string   `json:"path"`
	Params   []string `json:"params"`
	HostName string   `json:"hostname"`
	IPs      []string `json:"ips"`
}

// GetInfo returns the information about the current binary
func GetInfo() *Info {
	info := &Info{}

	// full path + binary name
	info.Path = os.Args[0]

	// if path, err := os.Getwd(); err == nil {
	// 	info.Path = os.Args[0] + path
	// }

	// params
	info.Params = os.Args[1:]

	// hostname
	if hostname, err := os.Hostname(); err == nil {
		info.HostName = hostname
	}

	// ips
	if ifaces, err := net.Interfaces(); err == nil {
		for _, i := range ifaces {
			if addrs, err := i.Addrs(); err == nil {
				for _, addr := range addrs {
					var ip net.IP

					switch v := addr.(type) {
					case *net.IPNet:
						ip = v.IP
					case *net.IPAddr:
						ip = v.IP
					}

					// IP addresses
					if !strings.Contains(ip.String(), ":") && ip.String() != "127.0.0.1" {
						info.IPs = append(info.IPs, ip.String())
					}
				}
			}
		}
	}

	return info
}
