package icons

import (
	"bytes"
	"strings"
	"sync"

	lucide "github.com/eduardolat/gomponents-lucide"
)

type Icon struct {
	lucide.IconInfo
	SearchTerms []string
	SVG         string
}

var (
	icons     []Icon
	iconsOnce sync.Once
)

func GetAllIcons() []Icon {
	iconsOnce.Do(func() {
		icons = make([]Icon, len(lucide.IconsInfo))

		for i, iconInfo := range lucide.IconsInfo {
			var svgBuffer bytes.Buffer
			iconNode := iconInfo.Icon()
			iconNode.Render(&svgBuffer)

			searchTerms := make([]string, 0, 1+len(iconInfo.Tags)+len(iconInfo.Categories))
			searchTerms = append(searchTerms, strings.ToLower(iconInfo.Name))
			searchTerms = append(searchTerms, lowercaseStrings(iconInfo.Tags)...)
			searchTerms = append(searchTerms, lowercaseStrings(iconInfo.Categories)...)

			icons[i] = Icon{
				IconInfo:    iconInfo,
				SearchTerms: searchTerms,
				SVG:         svgBuffer.String(),
			}
		}
	})

	return icons
}

func lowercaseStrings(strs []string) []string {
	lower := make([]string, len(strs))
	for i, s := range strs {
		lower[i] = strings.ToLower(s)
	}
	return lower
}
