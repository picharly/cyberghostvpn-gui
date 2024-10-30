package cg

import (
	"cyberghostvpn-gui/resources"
	"cyberghostvpn-gui/tools"
	"sort"
	"strconv"
	"strings"
)

var cities *[]resources.City

// GetCities retrieves a list of cities for the specified server type and country code.
// It loads the cities if they are not already loaded and returns a pointer to the list of cities.
func GetCities(serverType CgServerType, countryCode string) *[]resources.City {
	LoadCities(serverType, countryCode)

	return cities
}

// GetCity returns a City object given a city name.
// If the city is not found, it will return an empty City object.
func GetCity(name string) resources.City {
	if cities != nil {
		for _, c := range *cities {
			if c.Name == name {
				return c
			}
		}
	}
	return resources.City{}
}

// LoadCities loads the list of cities for the given server type and country code.
// It will store the result in the cities variable and return an error if the command fails.
// If the command succeeds, it will return nil.
func LoadCities(serverType CgServerType, countryCode string) error {
	array := make([]resources.City, 0)
	streaming := ""

	if serverType == CG_SERVER_TYPE_STREAMING && len(SelectedStreamingService) > 0 {
		streaming = `"` + SelectedStreamingService + `"`
	}

	if out, err := tools.ExecuteCommand(
		getCGCommand(
			GetOptionServerType(string(serverType)),
			streaming,
			string(CG_OTHER_COUNTRY_CODE),
			countryCode), true, false); err != nil {
		cities = &array
		return err
	} else {
		for _, line := range out {
			if strings.Contains(line, "|") {
				line = strings.Trim(line, "|")
				parts := strings.Split(line, "|")
				if len(parts) == 4 {
					c := resources.City{}
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
					if instance := strings.TrimSpace(parts[2]); len(instance) > 0 {
						c.Instance = instance
					}
					if load := strings.TrimSpace(parts[3]); len(load) > 0 {
						c.Load = load
					}
					array = append(array, c)
				}
			}
		}
		sort.Slice(array, func(i, j int) bool {
			return array[i].Name < array[j].Name
		})
		cities = &array
	}
	return nil
}
