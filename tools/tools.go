package tools

import (
	"bytes"
	"context"
	"fmt"
	"image/jpeg"
	"image/png"
	"log"
	"net/http"
	"os"
	"os/exec"
	"os/user"
	"strings"

	"github.com/shirou/gopsutil/v3/process"
)

// IsCommandExists checks if a command is available on the system.
//
// It does so by executing the "which" command and seeing if the command is
// available in the PATH. If it is, the function returns true. Otherwise, it
// returns false.
func IsCommandExists(cmd string) bool {
	_, err := exec.LookPath(cmd)
	return err == nil
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

// IsRoot returns true if the program is running as root user.
func IsRoot() bool {
	currentUser, err := user.Current()
	if err != nil {
		log.Fatalf("[isRoot] Unable to get current user: %s", err)
	}
	return currentUser.Username == "root"
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

// ToJpeg takes image bytes as input and returns the equivalent JPEG image.
//
// It currently supports PNG images, and will return an error if the input is not a PNG.
//
// The returned bytes are the JPEG image data.
func ToJpeg(imageBytes []byte) ([]byte, error) {

	// DetectContentType detects the content type
	contentType := http.DetectContentType(imageBytes)

	switch contentType {
	case "image/png":
		// Decode the PNG image bytes
		img, err := png.Decode(bytes.NewReader(imageBytes))

		if err != nil {
			return nil, err
		}

		buf := new(bytes.Buffer)

		// encode the image as a JPEG file
		if err := jpeg.Encode(buf, img, nil); err != nil {
			return nil, err
		}

		return buf.Bytes(), nil
	}

	return nil, fmt.Errorf("unable to convert %#v to jpeg", contentType)
}
