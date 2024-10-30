package cg

import (
	"cyberghostvpn-gui/resources"
	"cyberghostvpn-gui/settings"
)

var SelectedCountry resources.Country
var SelectedCity resources.City
var SelectedProtocol string
var SelectedServer resources.Server
var SelectedServiceType string
var SelectedStreamingService string
var SelectedVPNService string

// Server Type Options
var ServerTypeOptions = map[string]string{
	string(CG_SERVER_TYPE_TRAFFIC):   string(CG_SERVERTYPES_TRAFFIC),
	string(CG_SERVER_TYPE_STREAMING): string(CG_SERVERTYPES_STREAMING),
	string(CG_SERVER_TYPE_TORRENT):   string(CG_SERVERTYPES_TORRENT),
}

// VPN Service Options
var VPNServiceOptions = map[string]string{
	string(CG_SERVICE_TYPE_OPENVPN):   string(CG_SERVICES_OPENVPN),
	string(CG_SERVICE_TYPE_WIREGUARD): string(CG_SERVICES_WIREGUARD),
}

// GetOptionServerType returns the option for the given server type.
// It will return the option for traffic if the server type is unknown.
func GetOptionServerType(serverType string) string {
	for k, v := range ServerTypeOptions {
		if k == serverType {
			return v
		}
	}
	return string(CG_SERVER_TYPE_TRAFFIC)
}

// GetOptionVPNService returns the key for the given VPN service option.
// If the VPN service is not found, it defaults to returning the key for OpenVPN.
func GetOptionVPNService(vpnService string) string {
	for k, v := range VPNServiceOptions {
		if v == vpnService {
			return k
		}
	}
	return string(CG_SERVICE_TYPE_OPENVPN)
}

// DeleteProfile deletes the profile with the given name from the list of profiles and writes the updated
// list of profiles to the settings file.
func DeleteProfile(name string) {
	ps := settings.GetProfiles()
	for i, p := range *ps {
		if p.Name == name {
			*ps = append((*ps)[:i], (*ps)[i+1:]...)
			settings.WriteCurrentSettings()
			return
		}
	}
}

// SaveProfile saves the current profile with the given name to the list of profiles
// in the settings and writes the updated list of profiles to the settings file.
// If the previous name is given, it will update the existing profile with the given name
// instead of creating a new one.
func SaveProfile(name string, previousName string) {
	if len(name) > 0 {
		currentProfile := *settings.GetCurrentProfile()
		ps := settings.GetProfiles()
		slice := *ps

		// Update profile
		if len(previousName) > 0 {
			for i, p := range *ps {
				if p.Name == previousName {
					// Update profile
					p.Name = name
					p.City = currentProfile.City
					p.CountryCode = currentProfile.CountryCode
					p.CountryName = currentProfile.CountryName
					p.Protocol = currentProfile.Protocol
					p.Server = currentProfile.Server
					p.ServiceType = currentProfile.ServiceType
					p.StreamingService = currentProfile.StreamingService
					p.VPNService = currentProfile.VPNService

					// Update slice
					slice[i] = p
					*ps = slice
					settings.WriteCurrentSettings()
					return
				}
			}
		}

		// Or create a new one
		currentProfile.Name = name
		*ps = append(*ps, currentProfile)
		settings.WriteCurrentSettings()
	}
}

// SetSelectedCountry sets the currently selected country in the settings and writes the updated list of profiles to the settings file.
func SetSelectedCountry(country resources.Country) {
	SelectedCountry = country
	p := settings.GetCurrentProfile()
	p.CountryCode = country.Code
	p.CountryName = country.Name
	settings.WriteCurrentSettings()
}

// SetSelectedCity sets the currently selected city in the settings and writes the updated list of profiles to the settings file.
func SetSelectedCity(city resources.City) {
	SelectedCity = city
	p := settings.GetCurrentProfile()
	p.City = city.Name
	settings.WriteCurrentSettings()
}

// SetSelectedProtocol sets the currently selected protocol in the settings and writes the updated list of profiles to the settings file.
func SetSelectedProtocol(protocol string) {
	SelectedProtocol = protocol
	p := settings.GetCurrentProfile()
	p.Protocol = protocol
	settings.WriteCurrentSettings()
}

// SetSelectedServer sets the currently selected server in the settings and writes the updated list of profiles to the settings file.
func SetSelectedServer(server resources.Server) {
	SelectedServer = server
	p := settings.GetCurrentProfile()
	p.Server = server.Instance
	settings.WriteCurrentSettings()
}

// SetSelectedServiceType sets the currently selected service type in the settings and writes the updated list of profiles to the settings file.
func SetSelectedServiceType(serverType string) {
	SelectedServiceType = serverType
	p := settings.GetCurrentProfile()
	p.ServiceType = serverType
	settings.WriteCurrentSettings()
}

// SetSelectedVPNService sets the currently selected VPN service in the settings and writes the updated list of profiles to the settings file.
func SetSelectedVPNService(vpnService string) {
	GetOptionVPNService(vpnService)
	p := settings.GetCurrentProfile()
	p.VPNService = vpnService
	settings.WriteCurrentSettings()
}

// SetSelectedStreamingService sets the currently selected streaming service in the settings and writes the updated list of profiles to the settings file.
func SetSelectedStreamingService(streamingService string) {
	SelectedStreamingService = streamingService
	p := settings.GetCurrentProfile()
	p.StreamingService = streamingService
	settings.WriteCurrentSettings()
}

// GetServerType returns the CgServerType for the given server type name.
// If the name is not found in the ServerTypeOptions, it defaults to returning CG_SERVER_TYPE_TRAFFIC.
func GetServerType(name string) CgServerType {
	for k, _ := range ServerTypeOptions {
		if k == name {
			return CgServerType(k)
		}
	}
	return CG_SERVER_TYPE_TRAFFIC
}
