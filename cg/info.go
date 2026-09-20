package cg

import (
	"cyberghostvpn-gui/locales"
	"cyberghostvpn-gui/logger"
	"cyberghostvpn-gui/tools"
	"errors"
	"strings"
)

var CurrentState Status
var LastStatusError error
var LastStatusOutput string
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
		LastStatusError = nil
		LastStatusOutput = ""
	} else {
		newStatus, output, err := refreshStatus()
		switch newStatus {
		case string(cgConnected):
			CurrentState = Connected
		case string(cgNotConnected):
			CurrentState = Disconnected
		default:
			if err == nil {
				err = errors.New(locales.Text("err.inf3"))
			}
			if LastStatusError == nil {
				logger.Warnf("%s %v\n%s", locales.Text("err.inf2"), err, output)
			}
			CurrentState = Unknown
			LastStatusError = err
			LastStatusOutput = output
		}
		if CurrentState != Unknown {
			LastStatusError = nil
			LastStatusOutput = ""
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

// IsConnected returns true if the CyberGhost VPN client is currently connected, false otherwise.
func IsConnected() bool {
	return GetCurrentState() == Connected
}

// refreshStatus executes the CyberGhost VPN client with the status command and returns the first line of the output,
// the full command output, and an error if the command failed.
// If the command fails or the output is empty, it returns an empty status string.
func refreshStatus() (string, string, error) {
	// out, err := tools.ExecuteCommand(fmt.Sprintf("%s %s", CG_EXECUTABLE, CG_OTHER_STATUS), true, false)
	args := []string{
		string(CG_EXECUTABLE),
		string(CG_OTHER_STATUS),
	}
	out, err := tools.RunCommand(args, true, false, "")
	output := strings.Join(out, "\n")
	if err == nil && len(out) > 0 {
		return out[0], output, nil
	}
	return "", output, err
}
