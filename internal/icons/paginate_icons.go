package icons

// PaginateIcons returns a slice of icons for a given page, size and query.
// The query is used to filter the icons.
//
// The second return value is a boolean indicating if there are more icons to
// be paginated.
//
// The third return value is the next page of icons.
func PaginateIcons(page, size int, query string) ([]Icon, bool, int) {
	foundIcons := []Icon{}

	if query == "" {
		foundIcons = GetAllIcons()
	} else {
		foundIcons = SearchIcons(query)
	}

	start := (page - 1) * size
	if start >= len(foundIcons) {
		return []Icon{}, false, 1
	}

	end := start + size
	if end > len(foundIcons) {
		end = len(foundIcons)
	}

	hasNextPage := end < len(foundIcons)
	nextPage := page
	if hasNextPage {
		nextPage = page + 1
	}

	return foundIcons[start:end], hasNextPage, nextPage
}
