package cg

import (
	"cyberghostvpn-gui/locales"
	"cyberghostvpn-gui/logger"
	"cyberghostvpn-gui/tools"
	"strings"
)

var CurrentState Status
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
		CurrentState = NotInstalled
	} else {
		newStatus := refreshStatus()
		switch newStatus {
		case string(cgConnected):
			CurrentState = Connected
		case string(cgNotConnected):
			CurrentState = Disconnected
		default:
			CurrentState = Unknown
		}
	}

	return CurrentState
}

// GetVersion returns the version of the CyberGhost VPN client executable.
// If the executable command does not exist or the command fails, it will return an empty string.
func GetVersion() string {
	if len(version) < 1 {
		// out, err := tools.ExecuteCommand(fmt.Sprintf("%s %s | grep -i \"cyberghost -\"", CG_EXECUTABLE, CG_OTHER_HELP), true, false)
		args := []string{
			string(CG_EXECUTABLE),
			string(CG_OTHER_HELP),
		}
		out, err := tools.RunCommand(args, true, false, "")
		if err == nil && len(out) > 0 {
			for _, line := range out {
				if strings.Contains(line, "cyberghost -") {
					version = strings.ReplaceAll(strings.Replace(line, "cyberghost -", "", 1), " ", "")
					return version
				}
			}
		} else if err != nil {
			logger.Warnf("%s %sv", locales.Text("err.inf0"), err)
		}
	}
	return version
}

// refreshStatus executes the CyberGhost VPN client with the status command and returns the first line of the output.
// If the command fails or the output is empty, it returns an empty string.
func refreshStatus() string {
	// out, err := tools.ExecuteCommand(fmt.Sprintf("%s %s", CG_EXECUTABLE, CG_OTHER_STATUS), true, false)
	args := []string{
		string(CG_EXECUTABLE),
		string(CG_OTHER_STATUS),
	}
	out, err := tools.RunCommand(args, true, false, "")
	if err == nil && len(out) > 0 {
		return out[0]
	}
	return ""
}
