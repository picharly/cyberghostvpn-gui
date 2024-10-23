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

func GetCountries(serverType cgServerType) *[]resources.Country {
	if countries == nil || len(*countries) == 0 {
		LoadCountries(serverType)
	}

	return countries
}

func GetCountry(name string) resources.Country {
	for _, c := range *countries {
		if c.Name == name {
			return c
		}
	}
	return resources.Country{}
}

func LoadCountries(serverType cgServerType) error {
	array := make([]resources.Country, 0)

	if out, err := tools.ExecuteCommand(getCGCommand(getOptionServerType(serverType), string(CG_OTHER_COUNTRY_CODE)), true); err != nil {
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
