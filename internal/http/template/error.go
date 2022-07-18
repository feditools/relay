package template

// ErrorPageName is the name of the error template.
const ErrorName = "error"

// Error contains the variables for the error template.
type Error struct {
	Common

	Header      string
	Image       string
	SubHeader   string
	Paragraph   string
	ButtonHRef  string
	ButtonLabel string
}
