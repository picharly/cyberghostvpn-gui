package tools

/*
	Network-related function set
	2018-09: C.Pichard
*/

import (
	"cyberghostvpn-gui/locales"
	"fmt"
	"net"
)

// GetLocalIPAddresses returns a list of all local IP addresses.
//
// The function will return all IP addresses found on all local interfaces
// except for the local IP address (127.0.0.1) and the IPv6 equivalent ::1.
//
// The function will return an error if any error occurs while trying to read
// the network interfaces or their IP addresses. If the error is not nil, but
// at least one IP address has been found, the error will be set to nil.
func GetLocalIPAddresses(filters ...net.Flags) ([]*net.IP, error) {
	var ipAddr []*net.IP
	var lastError error

	// Get local interface
	ifaces, err := net.Interfaces()
	if err != nil {
		lastError = fmt.Errorf("%s: %v", locales.Text("err.too2"), err)
	}

	// Iterate network interfaces
	for _, i := range ifaces {

		// Filtering
		if len(filters) > 0 {
			filter := false
			for _, f := range filters {
				if i.Flags&f == 0 {
					filter = true
					break
				}
			}
			if filter {
				continue
			}
		}

		addrs, err := i.Addrs()
		if err != nil {
			lastError = fmt.Errorf("%s: %v", locales.Text("err.too3", locales.Variable{Name: "Interface", Value: i.Name}), err)
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
			if ip != nil && ip.IsLoopback() {
				continue
			}
			ipAddr = append(ipAddr, &ip)
		}
	}

	// Be sure to clean any errors if one or more IP address has been found
	if lastError != nil && len(ipAddr) > 0 {
		lastError = nil
	}

	return ipAddr, lastError
}
