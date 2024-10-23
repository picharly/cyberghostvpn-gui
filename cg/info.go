package cg

import (
	"cyberghostvpn-gui/tools"
	"fmt"
	"strings"
)

var currentState Status

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

func GetCurrentState() Status {

	if !tools.IsCommandExists(string(CG_EXECUTABLE)) {
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

func GetVersion() string {
	out, err := tools.ExecuteCommand(fmt.Sprintf("%s %s | grep -i \"cyberghost -\"", CG_EXECUTABLE, CG_OTHER_HELP), true)
	if err == nil && len(out) > 0 {
		return strings.ReplaceAll(strings.Replace(out[0], "cyberghost -", "", 1), " ", "")
	}
	return ""
}

func refreshStatus() string {
	out, err := tools.ExecuteCommand(fmt.Sprintf("%s %s", CG_EXECUTABLE, CG_OTHER_STATUS), true)
	if err == nil && len(out) > 0 {
		return out[0]
	}
	return ""
}
