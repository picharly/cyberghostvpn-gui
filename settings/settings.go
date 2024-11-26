package settings

import (
	"cyberghostvpn-gui/locales"
	"cyberghostvpn-gui/resources"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
)

const (
	settingsFileName   = "settings.conf"
	settingsFolderName = "cyberghostvpn-gui"
)

var settingsFilePath string

var currentSettings *settings // Instanciating new Settings
var noSettings = true

type settings struct {
	ConnectStartup    bool                `json:"connect_startup"`
	Countries         []resources.Country `json:"countries"`
	HideOnClose       bool                `json:"hide_on_close"`
	KeepPassMem       bool                `json:"keep_password_memory"`
	Language          string              `json:"language"`
	LastProfile       Profile             `json:"last_profile"`
	LoadLastProfile   bool                `json:"load_last_profile"`
	HideOnStart       bool                `json:"hide_on_startup"`
	HideWhenConnected bool                `json:"hide_when_connected"`
	StopVPNOnExit     bool                `json:"stop_vpn_on_exit"`
	TrayIcon          bool                `json:"tray_icon"`
	Profiles          []Profile           `json:"profiles"`
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
	// Check file
	setSettingsFile()

	// Ensure directory exists
	if err := os.MkdirAll(filepath.Dir(settingsFilePath), 0755); err != nil {
		return err
	}

	// Write settings
	output, err := json.MarshalIndent(currentSettings, "", " ")
	if err != nil {
		return err
	}
	err = os.WriteFile(settingsFilePath, output, os.FileMode(int(0777)))
	if err != nil {
		return fmt.Errorf("%s: %v", locales.Text("err.set0", locales.Variable{Name: "FileName", Value: settingsFilePath}), err)
	}
	return nil
}

// setSettingsFile sets the full path of the settings file according to the operating system.
// For Linux, it is $HOME/.config/cyberghostvpn-gui/settings.conf.
// For Windows, it is %APPDATA%\cyberghostvpn-gui\settings.conf.
// For macOS, it is $HOME/Library/Application Support/cyberghostvpn-gui/settings.conf.
// For other operating systems, it is ./settings.conf.
func setSettingsFile() {
	if len(settingsFilePath) == 0 {
		switch runtime.GOOS {
		case "linux":
			settingsFilePath = filepath.Join(os.Getenv("HOME"), ".config", settingsFolderName, settingsFileName)
		case "windows":
			settingsFilePath = filepath.Join(os.Getenv("APPDATA"), settingsFolderName, settingsFileName)
		case "darwin": // macOS
			settingsFilePath = filepath.Join(os.Getenv("HOME"), "Library", "Application Support", settingsFolderName, settingsFileName)
		default:
			settingsFilePath = "./" + settingsFileName
		}
	}
}

// isSettingsFileExists checks if settings file exists on the file system.
func isSettingsFileExists() bool {
	if _, err := os.Stat(settingsFilePath); err != nil {
		return false
	}
	return true
}

// Function for reading settings from file
func readSettings() error {
	// Check file
	setSettingsFile()

	// Read settings
	if currentSettings == nil {
		if !isSettingsFileExists() {
			currentSettings = new(settings)
			currentSettings.Language = locales.GetSystemLocale()
			if err := WriteCurrentSettings(); err != nil {
				return err
			}
			noSettings = false
		}
		file, err := os.Open(settingsFilePath)
		if err != nil {
			return fmt.Errorf("%s: %v", locales.Text("err.set1", locales.Variable{Name: "FileName", Value: settingsFilePath}), err)
		}
		decoder := json.NewDecoder(file)
		err = decoder.Decode(&currentSettings)
		if err != nil {
			return fmt.Errorf("%s: %v", locales.Text("err.set2", locales.Variable{Name: "FileName", Value: settingsFilePath}), err)
		}
		noSettings = false
	}

	return nil
}
