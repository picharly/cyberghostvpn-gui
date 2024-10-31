package tools

import (
	"bufio"
	"context"
	"cyberghostvpn-gui/locales"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"

	"github.com/shirou/gopsutil/v3/process"
)

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

// ExecuteCommand runs a command and optionally returns the output of the command
// as a string array.
//
// The command is executed with bash as the shell. The command can be a simple
// command or a shell script.
//
// If the getOutput parameter is true, the function returns the output of the
// command as a string array. If the getOutput parameter is false, the function
// returns an empty string array.
//
// The function returns an error if the command execution fails.
func ExecuteCommand(command string, getOutput bool, sudo bool) ([]string, error) {

	sudoCmd := ""
	if sudo {
		if path, ok := IsCommandExists("gksudo"); ok {
			sudoCmd = path
		} else if path, ok := IsCommandExists("pkexec"); ok {
			sudoCmd = path
		}
	}

	var cmd *exec.Cmd
	if len(sudoCmd) > 0 {
		sudoCmd += " --user " + os.Getenv("USER")
		command = "\"" + sudoCmd + " " + command + "\""
	} else {
		cmd = exec.Command("bash", "-c", command)
	}

	stdout, _ := cmd.StdoutPipe()
	cmd.Start()
	scanner := bufio.NewScanner(stdout)
	scanner.Split(bufio.ScanLines)
	result := []string{}
	for scanner.Scan() {
		m := scanner.Text()
		if getOutput {
			result = append(result, m)
		}
	}
	err := cmd.Wait()
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
