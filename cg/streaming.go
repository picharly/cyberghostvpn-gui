package cg

import (
	"cyberghostvpn-gui/resources"
	"cyberghostvpn-gui/tools"
	"sort"
	"strconv"
	"strings"
)

var streamingServices *[]resources.StreamingService

// GetStreamingService returns a StreamingService object given a service name.
// It searches through the loaded streaming services and matches the service name.
// If no matching service is found, it returns an empty StreamingService object.
func GetStreamingService(name string) resources.StreamingService {
	if streamingServices != nil {
		for _, c := range *streamingServices {
			if c.Service == name {
				return c
			}
		}
	}
	return resources.StreamingService{}
}

// GetStreamingServices retrieves the list of streaming services available for a given country code.
// The function uses loadStreamingServices to fetch the streaming services and updates the streamingServices variable.
// It returns a pointer to the list of StreamingService objects.
func GetStreamingServices(countryCode string) *[]resources.StreamingService {
	streamingServices = loadStreamingServices(countryCode)
	return streamingServices
}

// loadStreamingServices loads the list of streaming services available for a given country code.
// It executes a command to retrieve streaming service data, parses the output, and
// populates the streamingServices variable. The function also updates the current
// settings with the retrieved streaming services and returns an error if the
// command fails. If the command succeeds, it returns nil.
//
// Parameters:
//
//	countryCode: The country code to filter the streaming services.
//
// Returns:
//
//	A pointer to the list of StreamingService objects.
func loadStreamingServices(countryCode string) *[]resources.StreamingService {
	array := make([]resources.StreamingService, 0)

	if len(countryCode) > 0 {

		args := []string{
			string(CG_EXECUTABLE),
			string(CG_SERVERTYPES_STREAMING),
			string(CG_OTHER_COUNTRY_CODE),
			countryCode,
		}
		if out, err := tools.RunCommand(args, true, false, ""); err != nil {
			streamingServices = &array
			return &array
		} else {
			for _, line := range out {
				if strings.Contains(line, "|") {
					line = strings.Trim(line, "|")
					parts := strings.Split(line, "|")
					if len(parts) == 3 {
						c := resources.StreamingService{}
						if id, err := strconv.Atoi((strings.TrimSpace(parts[0]))); err != nil {
							continue
						} else {
							c.Id = id
						}
						if name := strings.TrimSpace(parts[1]); len(name) > 0 {
							c.Service = name
						} else {
							continue
						}
						if code := strings.TrimSpace(parts[2]); len(code) > 0 {
							c.CountryCode = code
						} else {
							continue
						}
						array = append(array, c)
					}
				}
			}
		}
		sort.Slice(array, func(i, j int) bool {
			return array[i].Service < array[j].Service
		})
	}
	streamingServices = &array

	return streamingServices
}
