package cg

import (
	"strings"
)

// Connect builds a command to connect to a CyberGhost VPN server and returns it as a slice of strings.
//
// The command is built based on the values of the following global variables:
//
// - SelectedServiceType: the type of server (traffic, streaming, torrent)
// - SelectedStreamingService: the name of the streaming service to use when SelectedServiceType is streaming
// - SelectedCountry: the country to connect to
// - SelectedCity: the city to connect to
// - SelectedServer: the server to connect to
// - SelectedProtocol: the protocol to use (UDP or TCP)
// - SelectedVPNService: the VPN service to use (OpenVPN or WireGuard)
//
// The command is built by calling getCGCommandWithArgs with the relevant options.
//
// The returned slice of strings can be used with the os/exec package to execute the command.
func Connect() []string {
	options := []string{}

	// Server Type
	if SelectedServiceType == string(CG_SERVER_TYPE_STREAMING) && len(SelectedStreamingService) > 0 {
		options = append(options, string(CG_SERVERTYPES_STREAMING), `"`+SelectedStreamingService+`"`)
	} else {
		options = append(options, GetOptionServerType(SelectedServiceType))
	}

	// Country
	if SelectedCountry.Code != "" {
		options = append(options, string(CG_OTHER_COUNTRY_CODE), SelectedCountry.Code)
	}

	// City
	if SelectedCity.Name != "" {
		options = append(options, string(CG_OTHER_CITY), SelectedCity.Name)
	}

	// Server
	if SelectedServer.Name != "" {
		options = append(options, string(CG_OTHER_SERVER), SelectedServer.Instance)
	}

	// Protocol
	if SelectedProtocol != "" {
		options = append(options, string(CG_OTHER_CONNECTION), strings.ToLower(SelectedProtocol))
	}

	// VPN Service
	switch SelectedVPNService {
	case string(CG_SERVICE_TYPE_OPENVPN):
		options = append(options, string(CG_SERVICES_OPENVPN))
	case string(CG_SERVICE_TYPE_WIREGUARD):
		options = append(options, string(CG_SERVICES_WIREGUARD))
	}

	// Connect
	options = append(options, string(CG_OTHER_CONNECT))

	// Build command with SUDO
	cmd := getCGCommandWithArgs(options...)

	return cmd
}

// Disconnect returns a command to disconnect the VPN connection.
func Disconnect() []string {
	return getCGCommandWithArgs(string(CG_OTHER_STOP))
}
