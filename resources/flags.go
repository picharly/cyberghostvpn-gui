package resources

import (
	"cyberghostvpn-gui/locales"
	"cyberghostvpn-gui/logger"
	"fmt"
	"io/fs"
	"path/filepath"
	"strings"
)

type flag struct {
	Code string
	Data []byte
}

var flags = []flag{}

// GetFlag returns the SVG data for the given country code.
// The data is loaded from the "flags" directory on first call.
// If the given code is not found, it will return nil.
func GetFlag(code string) []byte {
	if len(flags) == 0 {
		loadFlags()
	}
	for _, f := range flags {
		if strings.EqualFold(f.Code, code) {
			return f.Data
		}
	}
	return nil
}

// loadFlags loads the list of flags for all countries.
// It will read all files in the "flags" directory and store them in the flags variable.
// If a file cannot be read, it will log a warning and continue with the next file.
func loadFlags() {
	// Read all files in the "flags" directory
	flagDir := "flags"
	files, err := srcFlags.ReadDir(flagDir)
	if err != nil {
		logger.Warnf(fmt.Sprintf("%s %s", locales.Text("res.err0"), err.Error()))
	}

	// Iterate over the files and print their names
	flags = make([]flag, len(files))
	for _, file := range files {
		if !file.IsDir() {
			data, err := fs.ReadFile(srcFlags, filepath.Join(flagDir, file.Name()))
			if err != nil {
				logger.Warnf(fmt.Sprintf("%s %s", locales.Text("res.err1", locales.Variable{Name: "FileName", Value: file.Name()}), err.Error()))
				continue
			}
			f := flag{
				Code: strings.TrimSuffix(file.Name(), ".svg"),
				Data: data,
			}
			flags = append(flags, f)
		}
	}

}
