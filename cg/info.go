package cg

import (
	"cyberghostvpn-gui/tools"
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

	if !tools.IsCommandExists("cyberghostvpn") {
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
	out, err := tools.ExecuteCommand("cyberghostvpn --help | grep -i \"cyberghost -\"", true)
	if err == nil && len(out) > 0 {
		return strings.ReplaceAll(strings.Replace(out[0], "cyberghost -", "", 1), " ", "")
	}
	return ""
}

func refreshStatus() string {
	out, err := tools.ExecuteCommand("cyberghostvpn --status", true)
	if err == nil && len(out) > 0 {
		return out[0]
	}
	return ""
}
