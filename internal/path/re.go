package path

import (
	"fmt"
	"regexp"
)

const (
	// reToken matches a token.
	reToken = `[a-zA-Z0-9_]{16,}` //#nosec G101
)

var (
	// ReAppAdminHome matches the admin home page.
	ReAppAdminHome = regexp.MustCompile(fmt.Sprintf(`^?%s$`, AppAdminHome))
	// ReAppAdminInstances matches the admin instances page.
	ReAppAdminInstances = regexp.MustCompile(fmt.Sprintf(`^?%s$`, AppAdminInstances))

	// ReAppHome matches the Home page.
	ReAppHome = regexp.MustCompile(fmt.Sprintf(`^?%s$`, AppHome))
)
