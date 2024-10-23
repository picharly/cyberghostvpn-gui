package settings

type Profile struct {
	City             string `json:"city"`
	CountryCode      string `json:"country_code"`
	Name             string `json:"name"`
	ServiceType      string `json:"service_type"`
	StreamingService string `json:"streaming_service"`
	TCP              bool   `json:"tcp"`
	Torrent          bool   `json:"torrent"`
	Traffic          bool   `json:"traffic"`
	WireGuard        bool   `json:"wireguard"`
}

func GetProfiles() *[]Profile {
	if currentSettings.Profiles == nil {
		currentSettings.Profiles = make([]Profile, 0)
	}
	return &currentSettings.Profiles
}
