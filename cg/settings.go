package cg

import (
	"cyberghostvpn-gui/resources"
	"cyberghostvpn-gui/settings"
)

var SelectedCountry resources.Country
var SelectedCity resources.City
var SelectedProtocol string
var SelectedServer resources.Server
var SelectedServerType string
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

func GetOptionServerType(serverType string) string {
	for k, v := range ServerTypeOptions {
		if v == serverType {
			return k
		}
	}
	return string(CG_SERVER_TYPE_TRAFFIC)
}

func GetOptionVPNService(vpnService string) string {
	for k, v := range VPNServiceOptions {
		if v == vpnService {
			return k
		}
	}
	return string(CG_SERVICE_TYPE_OPENVPN)
}

func SaveProfile(name string) {

}

func SetSelectedCountry(country resources.Country) {
	SelectedCountry = country
	p := settings.GetCurrentProfile()
	p.CountryCode = country.Code
	settings.WriteCurrentSettings()
}
func SetSelectedCity(city resources.City) {
	SelectedCity = city
	p := settings.GetCurrentProfile()
	p.City = city.Name
	settings.WriteCurrentSettings()
}
func SetSelectedProtocol(protocol string) {
	SelectedProtocol = protocol
	p := settings.GetCurrentProfile()
	p.Protocol = protocol
	settings.WriteCurrentSettings()
}
func SetSelectedServer(server resources.Server) {
	SelectedServer = server
	p := settings.GetCurrentProfile()
	p.Server = server.Instance
	settings.WriteCurrentSettings()
}
func SetSelectedServerType(serverType string) {
	SelectedServerType = serverType
	p := settings.GetCurrentProfile()
	p.ServerType = serverType
	settings.WriteCurrentSettings()
}
func SetSelectedVPNService(vpnService string) {
	GetOptionVPNService(vpnService)
	p := settings.GetCurrentProfile()
	p.VPNService = vpnService
	settings.WriteCurrentSettings()
}

func GetServerType(name string) CgServerType {
	for k, _ := range ServerTypeOptions {
		if k == name {
			return CgServerType(k)
		}
	}
	return CG_SERVER_TYPE_TRAFFIC
}
