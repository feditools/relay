package template

import liblanguage "github.com/feditools/go-lib/language"

const (
	// FormInputTypeHidden is a hidden html input field.
	FormInputTypeHidden = "hidden"
	// FormInputTypePassword is a password html input field.
	FormInputTypePassword = "password"
	// FormInputTypeText is a text html input field.
	FormInputTypeText = "text"
)

// FormInput is a templated form input.
type FormInput struct {
	ID           string
	Type         string
	Name         string
	Placeholder  string
	Label        *liblanguage.LocalizedString
	LabelClass   string
	Value        string
	WrappedClass string
	Disabled     bool
	Required     bool
	Validation   *FormValidation
}

// FormValidation is a validation response to a form input.
type FormValidation struct {
	Valid    bool
	Response string
}
