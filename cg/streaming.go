package cg

import (
	"cyberghostvpn-gui/resources"
	"cyberghostvpn-gui/tools"
	"sort"
	"strconv"
	"strings"
)

var streamingServices *[]resources.StreamingService

func GetStreamingService(name string) resources.StreamingService {
	for _, c := range *streamingServices {
		if c.Service == name {
			return c
		}
	}
	return resources.StreamingService{}
}

func GetStreamingServices(countryCode string) *[]resources.StreamingService {
	streamingServices = loadStreamingServices(countryCode)
	return streamingServices
}

func loadStreamingServices(countryCode string) *[]resources.StreamingService {
	array := make([]resources.StreamingService, 0)

	if len(countryCode) > 0 {

		if out, err := tools.ExecuteCommand(getCGCommand(string(CG_SERVERTYPES_STREAMING), string(CG_OTHER_COUNTRY_CODE), countryCode), true); err != nil {
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
