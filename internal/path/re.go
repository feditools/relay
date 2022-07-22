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
	// ReAppAdminBlockPre matches the admin block page.
	ReAppAdminBlockPre = regexp.MustCompile(fmt.Sprintf(`^?%s`, AppAdminBlockList))
	// ReAppAdminInstancesPre matches the admin instances page.
	ReAppAdminInstancesPre = regexp.MustCompile(fmt.Sprintf(`^?%s`, AppAdminInstanceList))

	// ReAppHome matches the Home page.
	ReAppHome = regexp.MustCompile(fmt.Sprintf(`^?%s$`, AppHome))
)
