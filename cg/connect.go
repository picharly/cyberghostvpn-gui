package cg

import (
	"cyberghostvpn-gui/tools"
	"strings"
)

func Connect() (string, error) {
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

	// Execute command with SUDO
	// if out, err := tools.ExecuteCommand(getCGCommand(
	// 	options...), true, true); err != nil {
	// 	return strings.Join(out, "\n"), err
	// }

	return "", nil
}

func Disconnect() (string, error) {
	return tools.RunCommandWithGksudo(getCGCommand(
		string(CG_OTHER_STOP)))
}
