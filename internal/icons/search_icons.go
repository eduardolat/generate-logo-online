package icons

import (
	"strings"
)

func SearchIcons(query string) []Icon {
	allIcons := GetAllIcons()
	query = strings.ToLower(query)

	var result []Icon

	for _, icon := range allIcons {
		if matchIcon(icon, query) {
			result = append(result, icon)
		}
	}

	return result
}

func matchIcon(icon Icon, query string) bool {
	for _, keyword := range icon.SearchTerms {
		if strings.Contains(keyword, query) {
			return true
		}
	}
	return false
}
