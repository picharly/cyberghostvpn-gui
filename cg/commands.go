package cg

import (
	"fmt"
	"os/exec"
	"strings"
)

type cgCommand string

const (
	CG_EXECUTABLE cgCommand = "cyberghostvpn"

	CG_SERVICES_TYPE      cgCommand = "--service-type" // needs a service name
	CG_SERVICES_OPENVPN   cgCommand = "--openvpn"
	CG_SERVICES_WIREGUARD cgCommand = "--wireguard"

	CG_SERVERTYPES_TYPE      cgCommand = "--server-type" // needs a server type (traffic, streaming, torrent)
	CG_SERVERTYPES_TRAFFIC   cgCommand = "--traffic"
	CG_SERVERTYPES_STREAMING cgCommand = "--streaming" // needs a streaming service name
	CG_SERVERTYPES_TORRENT   cgCommand = "--torrent"

	CG_OTHER_COUNTRY_CODE cgCommand = "--country-code" // needs a country code
	CG_OTHER_CONNECTION   cgCommand = "--connection"   // needs a connection type (UDP, TCP)
	CG_OTHER_CITY         cgCommand = "--city"         // needs a city name
	CG_OTHER_SERVER       cgCommand = "--server"       // needs a server name
	CG_OTHER_CONNECT      cgCommand = "--connect"
	CG_OTHER_HELP         cgCommand = "--help"
	CG_OTHER_STATUS       cgCommand = "--status"
	CG_OTHER_STOP         cgCommand = "--stop"
	CG_OTHER_SETUP        cgCommand = "--setup"
	CG_OTHER_UNINSTALL    cgCommand = "--uninstall"
)

type cgServiceType string

const (
	CG_SERVICE_TYPE_OPENVPN   cgServiceType = "OpenVPN"
	CG_SERVICE_TYPE_WIREGUARD cgServiceType = "WireGuard"
)

type CgServerType string

const (
	CG_SERVER_TYPE_TRAFFIC   CgServerType = "Traffic"
	CG_SERVER_TYPE_STREAMING CgServerType = "Streaming"
	CG_SERVER_TYPE_TORRENT   CgServerType = "Torrent"
)

type cgConnection string

const (
	CG_CONNECTION_UDP cgConnection = "UDP"
	CG_CONNECTION_TCP cgConnection = "TCP"
)

// getCGCommand constructs a command string for executing the CyberGhost VPN CLI.
//
// It takes a variable number of string options, joins them with spaces, and
// appends them to the base command defined by CG_EXECUTABLE.
//
// The constructed command string is printed to the standard output for
// debugging purposes, and then returned.
//
// Parameters:
//
//	options: a variable number of string arguments to append as options
//	         to the base command.
//
// Returns:
//
//	A string representing the complete command to be executed.
func getCGCommand(options ...string) string {

	// Look for path
	path, err := exec.LookPath(string(CG_EXECUTABLE))
	if err != nil {
		path = string(CG_EXECUTABLE)
	}

	cmd := path + " "
	if len(options) > 0 {
		cmd += strings.Join(options, " ")
	}
	fmt.Printf("Command: %s\n", cmd)

	return cmd
}
