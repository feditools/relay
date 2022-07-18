package http

import (
	"net/url"
	"strconv"
)

// GetPaginationFromURL returns a page and item count derived from the url query returns true if count is present in query.
func GetPaginationFromURL(u *url.URL, defaultCount int) (page int, count int, countFound bool) {
	// get display count
	iCount := defaultCount
	if qCount, ok := u.Query()["count"]; ok {
		if len(qCount[0]) >= 1 {
			uCount, err := strconv.ParseUint(qCount[0], 10, 64)
			if err == nil {
				iCount = int(uCount)
			}
		}
	}

	// get display page
	iPage := 1
	if qPage, ok := u.Query()["page"]; ok {
		if len(qPage[0]) >= 1 {
			uPage, err := strconv.ParseUint(qPage[0], 10, 64)
			if err == nil {
				iPage = int(uPage)
			}
		}
	}

	return iPage, iCount, defaultCount != iCount
}
