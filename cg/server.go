package cg

import (
	"cyberghostvpn-gui/resources"
	"cyberghostvpn-gui/tools"
	"sort"
	"strconv"
	"strings"
)

var servers *[]resources.Server

// GetServers retrieves a list of servers for the specified server type, country code and city name.
// It loads the servers if they are not already loaded and returns a pointer to the list of servers.
func GetServers(serverType CgServerType, countryCode string, cityName string) *[]resources.Server {
	LoadServers(serverType, countryCode, cityName)

	return servers
}

// GetServer returns a Server object given an instance name.
// It searches through the loaded servers and matches the instance name.
// If the instance name contains a space, it considers only the first part.
// If no matching server is found, it returns an empty Server object.
func GetServer(instance string) resources.Server {
	if servers != nil {
		for _, s := range *servers {
			if strings.Contains(strings.ToLower(instance), " ") {
				split := strings.Split(instance, " ")
				instance = split[0]
			}
			if s.Instance == instance {
				return s
			}
		}
	}
	return resources.Server{}
}

// LoadServers loads the list of servers for the given server type, country code and city name.
// It executes a command to retrieve server data, parses the output, and
// populates the servers variable. The function also updates the current
// settings with the retrieved server list and returns an error if the
// command fails. If the command succeeds, it returns nil.
//
// Parameters:
//
//	serverType: The server type to filter the servers.
//	countryCode: The country code to filter the servers.
//	cityName: The city name to filter the servers.
//
// Returns:
//
//	An error if the command execution fails, or nil if it succeeds.
func LoadServers(serverType CgServerType, countryCode string, cityName string) error {
	array := make([]resources.Server, 0)
	streaming := ""

	if serverType == CG_SERVER_TYPE_STREAMING && len(SelectedStreamingService) > 0 {
		streaming = `"` + SelectedStreamingService + `"`
	}

	args := []string{
		string(CG_EXECUTABLE),
		GetOptionServerType(string(serverType)),
		streaming,
		string(CG_OTHER_COUNTRY_CODE),
		countryCode,
		string(CG_OTHER_CITY),
		cityName,
	}
	if out, err := tools.RunCommand(args, true, false, ""); err != nil {
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
