package tools

import (
	"context"
	"cyberghostvpn-gui/locales"
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"runtime"
	"strings"

	"github.com/shirou/gopsutil/v3/process"
)

var NeedsPassword io.WriteCloser
var PasswordChannel = make(chan string)

// IsCommandExists checks if a command exists on the system.
//
// It does so by using the exec.LookPath function. If the command does not exist,
// the function returns false. Otherwise, it returns true. Additionally, the
// function returns the path of the command if it was found, or the original
// command name if it was not.
func IsCommandExists(cmd string) (string, bool) {
	path, err := exec.LookPath(cmd)
	if len(path) < 1 {
		path = cmd
	}
	return path, err == nil
}

// IsFileExists checks if a file exists on the file system.
//
// It does so by using the os.Stat function. If the file does not exist,
// the function returns false. Otherwise, it returns true.
func IsFileExists(name string) bool {
	if len(name) > 0 {
		if _, err := os.Stat(name); os.IsNotExist(err) {
			return false
		} else {
			return true
		}
	}
	return false
}

// IsServiceRunning checks if a service is running.
//
// It does so by getting the list of all running processes and
// checking if the given service name is in that list. If it is,
// the function returns true. Otherwise, it returns false.
func IsServiceRunning(name string) bool {
	if p, err := process.Processes(); err == nil && len(p) > 0 {
		for i := 0; i < len(p); i++ {
			proc := p[i]
			if procName, err := proc.NameWithContext(context.Background()); err == nil && len(procName) > 0 {
				if strings.Contains(strings.ToLower(name), strings.ToLower(procName)) {
					return true
				}
			}

		}
	}
	return false
}

// RunCommand runs a command with arguments and returns the output as a string array.
//
// Parameters:
//
//	args: a slice of strings representing the command and its arguments
//	getOutput: boolean indicating whether the function should return the output of the command
//	sudo: boolean indicating whether the command should be run with sudo
//	pwd: string containing the password to use with sudo
//
// Returns:
//
//	A string array containing the output of the command if getOutput is true, an empty string array otherwise
//	An error if the command execution fails
func RunCommand(args []string, getOutput bool, sudo bool, pwd string) ([]string, error) {

	// Add sudo command if needed
	if sudo {
		newArgs := make([]string, 0)
		newArgs = append(newArgs, "sudo", "-S", "--")
		newArgs = append(newArgs, args...)
		args = newArgs
	}

	// Instance the command
	cmd := exec.Command(args[0], args[1:]...)
	if sudo {
		cmd.Stdin = strings.NewReader(pwd)
	}

	// Start the command
	stdout, err := cmd.CombinedOutput()
	result := strings.Split(string(stdout), "\n")

	return result, err
}

// RunCommandWithGksudo runs a command with gksudo or pkexec if gksudo is not available.
//
// It currently supports Linux systems only.
//
// The function takes a command and a variable number of arguments as input.
// The command is executed with gksudo if available, otherwise with pkexec.
// The output of the command is returned as a string.
//
// If an error occurs during the execution of the command, the function returns
// an error.
//
// NOT USED BECAUSE CYBERGHOSTVPN CLI ASK FOR SUDO, IT DOES NOT WORK WITH PKEXEC
func RunCommandWithGksudo(command string, args ...string) (string, error) {
	var sudoCmd *exec.Cmd

	// Use gksudo on Linux, pkexec if gksudo is not available
	if runtime.GOOS == "linux" {
		_, err := exec.LookPath("gksudo")
		if err == nil {
			sudoCmd = exec.Command("gksudo", append([]string{command}, args...)...)
		} else {
			sudoCmd = exec.Command("pkexec", append([]string{command}, args...)...)
		}
	} else {
		return "", errors.New(locales.Text("err.too0"))
	}

	output, err := sudoCmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("%s: %w", locales.Text("err.too1"), err)
	}

	return string(output), nil
}
