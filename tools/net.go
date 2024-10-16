package tools

/*
	Network-related function set
	2018-09: C.Pichard
*/

import (
	"fmt"
	"net"
	"strings"
)

// GetLocalIPAddresses returns a list of all local IP addresses.
//
// The function will return all IP addresses found on all local interfaces
// except for the local IP address (127.0.0.1) and the IPv6 equivalent ::1.
//
// The function will return an error if any error occurs while trying to read
// the network interfaces or their IP addresses. If the error is not nil, but
// at least one IP address has been found, the error will be set to nil.
func GetLocalIPAddresses() ([]*net.IP, error) {
	var ipAddr []*net.IP
	var lastError error

	// Get local interface
	ifaces, err := net.Interfaces()
	if err != nil {
		lastError = fmt.Errorf("failed to retrieve network interfaces: %v", err)
	}

	// Iterate network interfaces
	for _, i := range ifaces {
		addrs, err := i.Addrs()
		if err != nil {
			lastError = fmt.Errorf("failed to read ip addresse from interface '%s': %v", i.Name, err)
			continue
		}
		// Iterate IP addresses
		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}
			// Ignore local IP address
			if ip != nil && !strings.HasPrefix(ip.String(), "127") && !strings.EqualFold(ip.String(), "::1") {
				ipAddr = append(ipAddr, &ip)
			}
		}
	}

	// Be sure to clean any errors if one or more IP address has been found
	if lastError != nil && len(ipAddr) > 0 {
		lastError = nil
	}

	return ipAddr, lastError
}
