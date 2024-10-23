package settings

import (
	"cyberghostvpn-gui/locales"
	"cyberghostvpn-gui/resources"
	"encoding/json"
	"fmt"
	"os"
)

const settingsFile string = "./settings.conf"

var currentSettings *settings // Instanciating new Settings
var noSettings = true

type settings struct {
	Countries   []resources.Country `json:"countries"`
	HideOnClose bool                `json:"hide_on_close"`
	Language    string              `json:"language"`
	LastProfile []Profile           `json:"last_profile"`
	HideOnStart bool                `json:"hide_on_startup"`
	TrayIcon    bool                `json:"tray_icon"`
	Profiles    []Profile           `json:"profiles"`
}

// Get current settings
func GetCurrentSettings() (*settings, error) {
	err := readSettings()
	return currentSettings, err
}

// Check if settings are ready
func IsSettingsOK() bool {
	err := readSettings()
	return err == nil
}

// Create new settings instance
// /!\ WILL ERASE CURRENT SETTINGS /!\
func NewSettings() *settings {
	currentSettings = &settings{}
	return currentSettings
}

// Writing settings file
func WriteCurrentSettings() error {
	output, err := json.MarshalIndent(currentSettings, "", " ")
	if err != nil {
		return err
	}
	err = os.WriteFile(settingsFile, output, os.FileMode(int(0777)))
	if err != nil {
		return fmt.Errorf("%s: %v", locales.Text("err.set0", locales.Variable{Name: "FileName", Value: settingsFile}), err)
	}
	return nil
}

// isSettingsFileExists checks if settings file exists on the file system.
func isSettingsFileExists() bool {
	if _, err := os.Stat(settingsFile); err != nil {
		return false
	}
	return true
}

// Function for reading settings from file
func readSettings() error {
	if currentSettings == nil {
		if !isSettingsFileExists() {
			currentSettings = new(settings)
			currentSettings.Language = locales.GetSystemLocale()
			if err := WriteCurrentSettings(); err != nil {
				return err
			}
			noSettings = false
		}
		file, err := os.Open(settingsFile)
		if err != nil {
			return fmt.Errorf("%s: %v", locales.Text("err.set1", locales.Variable{Name: "FileName", Value: settingsFile}), err)
		}
		decoder := json.NewDecoder(file)
		err = decoder.Decode(&currentSettings)
		if err != nil {
			return fmt.Errorf("%s: %v", locales.Text("err.set2", locales.Variable{Name: "FileName", Value: settingsFile}), err)
		}
		noSettings = false
	}

	return nil
}
