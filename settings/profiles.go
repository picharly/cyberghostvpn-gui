package settings

type Profile struct {
	City             string `json:"city"`
	CountryCode      string `json:"country_code"`
	CountryName      string `json:"country_name"`
	Name             string `json:"name"`
	Protocol         string `json:"protocol"`
	Server           string `json:"server"`
	ServiceType      string `json:"service_type"`
	StreamingService string `json:"streaming_service"`
	VPNService       string `json:"vpn_service"`
}

func GetCurrentProfile() *Profile {
	return &currentSettings.LastProfile
}

func GetProfile(name string) *Profile {
	for _, p := range *GetProfiles() {
		if p.Name == name {
			return &p
		}
	}
	return nil
}

func GetProfiles() *[]Profile {
	if currentSettings.Profiles == nil {
		currentSettings.Profiles = make([]Profile, 0)
	}
	return &currentSettings.Profiles
}
