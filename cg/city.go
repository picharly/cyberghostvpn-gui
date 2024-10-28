package cg

import (
	"cyberghostvpn-gui/resources"
	"cyberghostvpn-gui/tools"
	"sort"
	"strconv"
	"strings"
)

var cities *[]resources.City

func GetCities(serverType CgServerType, countryCode string) *[]resources.City {
	LoadCities(serverType, countryCode)

	return cities
}

func GetCity(name string) resources.City {
	for _, c := range *cities {
		if c.Name == name {
			return c
		}
	}
	return resources.City{}
}

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
			countryCode), true); err != nil {
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
