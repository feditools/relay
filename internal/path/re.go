package path

import (
	"fmt"
	"regexp"
)

var (
	// ReAdmin matches the admin page.
	ReAdmin = regexp.MustCompile(fmt.Sprintf(`^?/%s$`, PartAdmin))

	// ReHome matches the Home page.
	ReHome = regexp.MustCompile(fmt.Sprintf(`^%s$`, Home))
)
