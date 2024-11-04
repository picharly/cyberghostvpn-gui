package settings

import "cyberghostvpn-gui/tools"

var requirements = []string{"cyberghostvpn", "sudo"}

// CheckRequirements returns a list of requirements that are not satisfied. The second argument
// is a boolean indicating whether all requirements are satisfied.
func CheckRequirements() ([]string, bool) {
	missing := []string{}
	for _, req := range requirements {
		if _, ok := tools.IsCommandExists(req); !ok {
			missing = append(missing, req)
		}
	}
	return missing, len(missing) == 0
}
