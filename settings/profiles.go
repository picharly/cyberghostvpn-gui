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

// GetCurrentProfile returns the last profile used.
func GetCurrentProfile() *Profile {
	return &currentSettings.LastProfile
}

// GetProfile returns the profile with the given name.
// If the profile is not found, it will return nil.
func GetProfile(name string) *Profile {
	for _, p := range *GetProfiles() {
		if p.Name == name {
			return &p
		}
	}
	return nil
}

// GetProfiles returns a pointer to the list of profiles stored in the settings.
// If the list of profiles is empty, it will initialize it with an empty list.
func GetProfiles() *[]Profile {
	if currentSettings.Profiles == nil {
		currentSettings.Profiles = make([]Profile, 0)
	}
	return &(*currentSettings).Profiles
}
