package template

import (
	"fmt"
	"math"
)

// Pagination is a pagination element that can be added to a webpage.
type Pagination []PaginationNode

// PaginationNode is an element in a pagination element.
type PaginationNode struct {
	Text string
	Icon string
	HRef string

	Active   bool
	Disabled bool
}

// PaginationConfig contains the config to construct pagination.
type PaginationConfig struct {
	Count         int    // item count
	DisplayCount  int    // how many items to display per page
	HRef          string // href to add query to
	HRefCount     int    // count to include in the href, if 0 no count is added
	MaxPagination int    // the max number of pages to show
	Page          int    // current page
}

// MakePagination creates a pagination element from the provided parameters.
func MakePagination(c *PaginationConfig) Pagination {
	displayItems := c.MaxPagination
	pages := int(math.Ceil(float64(c.Count) / float64(c.DisplayCount)))
	startingNumber := 1

	switch {
	case pages < displayItems:
		// less than
		displayItems = pages
	case c.Page > pages-displayItems/2:
		// end of the
		startingNumber = pages - displayItems + 1
	case c.Page > displayItems/2:
		// center active
		startingNumber = c.Page - displayItems/2
	}

	var items Pagination

	// previous button
	prevItem := PaginationNode{
		Text: "Previous",
		Icon: "caret-left",
	}
	switch {
	case c.Page == 1:
		prevItem.Disabled = true
	case c.HRefCount > 0:
		prevItem.HRef = fmt.Sprintf("%s?page=%d&count=%d", c.HRef, c.Page-1, c.HRefCount)
	default:
		prevItem.HRef = fmt.Sprintf("%s?page=%d", c.HRef, c.Page-1)
	}
	items = append(items, prevItem)

	// add pages
	for i := 0; i < displayItems; i++ {
		newItem := PaginationNode{
			Text: fmt.Sprintf("%d", startingNumber+i),
		}

		switch {
		case c.Page == startingNumber+i:
			newItem.Active = true
		case c.HRefCount > 0:
			newItem.HRef = fmt.Sprintf("%s?page=%d&count=%d", c.HRef, startingNumber+i, c.HRefCount)
		default:
			newItem.HRef = fmt.Sprintf("%s?page=%d", c.HRef, startingNumber+i)
		}

		items = append(items, newItem)
	}

	// next button
	nextItem := PaginationNode{
		Text: "Next",
		Icon: "caret-right",
	}
	switch {
	case c.Page == pages:
		nextItem.Disabled = true
	case c.HRefCount > 0:
		nextItem.HRef = fmt.Sprintf("%s?page=%d&count=%d", c.HRef, c.Page+1, c.HRefCount)
	default:
		nextItem.HRef = fmt.Sprintf("%s?page=%d", c.HRef, c.Page+1)
	}
	items = append(items, nextItem)

	return items
}
