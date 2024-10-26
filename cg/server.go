package cg

import (
	"cyberghostvpn-gui/resources"
	"cyberghostvpn-gui/tools"
	"sort"
	"strconv"
	"strings"
)

var servers *[]resources.Server

func GetServers(serverType CgServerType, countryCode string, cityName string) *[]resources.Server {
	LoadServers(serverType, countryCode, cityName)

	return servers
}

func GetServer(instance string) resources.Server {
	for _, s := range *servers {
		if strings.Contains(strings.ToLower(instance), " ") {
			split := strings.Split(instance, " ")
			instance = split[0]
		}
		if s.Instance == instance {
			return s
		}
	}
	return resources.Server{}
}

func LoadServers(serverType CgServerType, countryCode string, cityName string) error {
	array := make([]resources.Server, 0)

	if out, err := tools.ExecuteCommand(
		getCGCommand(
			GetOptionServerType(string(serverType)),
			string(CG_OTHER_COUNTRY_CODE),
			countryCode,
			string(CG_OTHER_CITY),
			cityName), true); err != nil {
		servers = &array
		return err
	} else {
		for _, line := range out {
			if strings.Contains(line, "|") {
				line = strings.Trim(line, "|")
				parts := strings.Split(line, "|")
				if len(parts) == 4 {
					c := resources.Server{}
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
			return array[i].Load < array[j].Load
		})
		servers = &array
	}
	return nil
}
