package cg

import (
	"cyberghostvpn-gui/resources"
	"cyberghostvpn-gui/settings"
	"cyberghostvpn-gui/tools"
	"sort"
	"strconv"
	"strings"
)

var countries *[]resources.Country

// GetCountries retrieves a list of countries for the specified server type.
// It loads the countries if they are not already loaded and returns a pointer
// to the list of countries.
func GetCountries(serverType CgServerType) *[]resources.Country {
	if countries == nil || len(*countries) == 0 {
		LoadCountries(serverType)
	}

	return countries
}

// GetCountry returns a Country object given a country name.
// If the country is not found, it will return an empty Country object.
func GetCountry(name string) resources.Country {
	if countries != nil {
		for _, c := range *countries {
			if c.Name == name {
				return c
			}
		}
	}
	return resources.Country{}
}

// LoadCountries loads the list of countries for the given server type.
// It executes a command to retrieve country data, parses the output, and
// populates the countries variable. The function also updates the current
// settings with the retrieved country list and returns an error if the
// command fails. If the command succeeds, it returns nil.
//
// Parameters:
//
//	serverType: The server type to filter the countries.
//
// Returns:
//
//	An error if the command execution fails, or nil if it succeeds.
func LoadCountries(serverType CgServerType) error {
	array := make([]resources.Country, 0)

	args := []string{
		string(CG_EXECUTABLE),
		GetOptionServerType(string(serverType)),
		string(CG_OTHER_COUNTRY_CODE),
	}
	if out, err := tools.RunCommand(args, true, false, ""); err != nil {
		countries = &array
		return err
	} else {
		for _, line := range out {
			if strings.Contains(line, "|") {
				line = strings.Trim(line, "|")
				parts := strings.Split(line, "|")
				if len(parts) == 3 {
					c := resources.Country{}
					if id, err := strconv.Atoi((strings.TrimSpace(parts[0]))); err != nil {
						continue
					} else {
						c.Id = id
					}
					if name := strings.TrimSpace(parts[1]); len(name) > 0 {
						c.Name = name
					} else {
						continue
					}
					if code := strings.TrimSpace(parts[2]); len(code) > 0 {
						c.Code = code
					} else {
						continue
					}
					array = append(array, c)
				}
			}
		}
		sort.Slice(array, func(i, j int) bool {
			return array[i].Name < array[j].Name
		})
		countries = &array

		// Update settings or read from them
		cfg, cfgErr := settings.GetCurrentSettings()
		if len(array) > 0 {
			if cfgErr == nil {
				cfg.Countries = array
				settings.WriteCurrentSettings()
			}
		} else if cfgErr == nil && len(cfg.Countries) > 0 {
			countries = &cfg.Countries
		}

	}
	return nil
}
