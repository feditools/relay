package template

import libtemplate "github.com/feditools/go-lib/template"

// Admin contains the variables used in nearly every admin template.
type Admin struct {
	Sidebar libtemplate.Sidebar
}
