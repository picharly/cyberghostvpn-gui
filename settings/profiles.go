package settings

type Profile struct {
	City        string `json:"city"`
	CountryCode string `json:"country_code"`
	Name        string `json:"name"`
	Protocol    string `json:"protocol"`
	Server      string `json:"server"`
	ServiceType string `json:"service_type"`
	VPNService  string `json:"vpn_service"`
}

func GetCurrentProfile() *Profile {
	return &currentSettings.LastProfile
}

func GetProfiles() *[]Profile {
	if currentSettings.Profiles == nil {
		currentSettings.Profiles = make([]Profile, 0)
	}
	return &currentSettings.Profiles
}
