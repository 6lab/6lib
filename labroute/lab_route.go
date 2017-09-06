package labroute

import (
	"log"
	"net"
	"net/http"
	"os"
)

// Get the local IP address of the current computer
func GetLocalIPAddress() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		os.Stderr.WriteString(err.Error() + "\n")
		os.Exit(1)
	}

	for _, a := range addrs {
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}

	return ""
}

// Get the Remote IP Address
func UserIP(r *http.Request) string {
	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		log.Printf("userip: %q is not IP:port", r.RemoteAddr)
		return ""
	}

	// Check ip validity
	userIP := net.ParseIP(ip)
	if userIP == nil {
		log.Printf("userip: %q is not IP:port", r.RemoteAddr)
		return ""
	}

	// Replace the short ::1 by his name
	if ip == "::1" {
		ip = "localhost"
	}

	return ip
}
