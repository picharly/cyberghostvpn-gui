package cg

import (
	"cyberghostvpn-gui/tools"
	"strings"
)

// Connect returns a command string to connect to the VPN using the selected options.
//
// This function takes into account the selected server type, country, city, server, protocol, and VPN service.
//
// The command string is constructed by appending the selected options to the base command defined by CG_EXECUTABLE.
//
// The constructed command string is then prepended with SUDO to ensure that the command is run with root privileges.
//
// Parameters: None
//
// Returns: A string representing the complete command to be executed.
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
		options = append(options, string(CG_OTHER_SERVER), SelectedServer.Name)
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

// Disconnect stops the current VPN connection.
//
// It returns the output of the command if it succeeds, or an error if it fails.
func Disconnect() (string, error) {
	out, err := tools.RunCommand(getCGCommandWithArgs(
		string(CG_OTHER_STOP)), true, true)
	if err != nil {
		return strings.Join(out, "\n"), err
	}
	return "", nil
}
