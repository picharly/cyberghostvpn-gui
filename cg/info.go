package cg

import (
	"cyberghostvpn-gui/tools"
	"fmt"
	"strings"
)

var currentState Status
var version string

type Status int

const (
	Unknown Status = iota
	Connected
	Disconnected
	Connecting
	Disconnecting
	NotInstalled
)

type cgMessage string

const (
	cgConnected    cgMessage = "VPN connection found."
	cgNotConnected cgMessage = "No VPN connections found."
)

// GetCurrentState returns the current state of the CyberGhost VPN client.
// It will use the executable command to check the status of the VPN.
// If the executable command does not exist, it will return NotInstalled.
// If the command exists, it will return Connected if the VPN is connected,
// Disconnected if it is disconnected, or Unknown if the status is unknown.
func GetCurrentState() Status {

	if _, ok := tools.IsCommandExists(string(CG_EXECUTABLE)); !ok {
		currentState = NotInstalled
	} else {
		newStatus := refreshStatus()
		switch newStatus {
		case string(cgConnected):
			currentState = Connected
		case string(cgNotConnected):
			currentState = Disconnected
		default:
			currentState = Unknown
		}
	}

	return currentState
}

// GetVersion returns the version of the CyberGhost VPN client executable.
// If the executable command does not exist or the command fails, it will return an empty string.
func GetVersion() string {
	if version == "" {
		out, err := tools.ExecuteCommand(fmt.Sprintf("%s %s | grep -i \"cyberghost -\"", CG_EXECUTABLE, CG_OTHER_HELP), true, false)
		if err == nil && len(out) > 0 {
			return strings.ReplaceAll(strings.Replace(out[0], "cyberghost -", "", 1), " ", "")
		}
	}
	return ""
}

// refreshStatus executes the CyberGhost VPN client with the status command and returns the first line of the output.
// If the command fails or the output is empty, it returns an empty string.
func refreshStatus() string {
	out, err := tools.ExecuteCommand(fmt.Sprintf("%s %s", CG_EXECUTABLE, CG_OTHER_STATUS), true, false)
	if err == nil && len(out) > 0 {
		return out[0]
	}
	return ""
}
