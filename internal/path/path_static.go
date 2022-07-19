package path

const (
	// files.

	// FileDefaultCSS is the css document applies to all pages.
	FileDefaultCSS = StaticCSS + "/default.min.css"
	// FileErrorCSS is the css document applies to the error page.
	FileErrorCSS = StaticCSS + "/error.min.css"
	// FileLoginCSS is the css document applies to the login page.
	FileLoginCSS = StaticCSS + "/login.min.css"

	// paths.

	// Static is the path for static files.
	Static = "/" + PartStatic + "/"
	// StaticCSS is the path.
	StaticCSS = Static + "css"
)
