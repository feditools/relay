package path

import (
	"fmt"
	"regexp"
)

var (
	// ReAppAdminHome matches the admin page.
	ReAppAdminHome = regexp.MustCompile(fmt.Sprintf(`^?%s$`, AppAdminHome))

	// ReAppHome matches the Home page.
	ReAppHome = regexp.MustCompile(fmt.Sprintf(`^?%s$`, AppHome))
)
