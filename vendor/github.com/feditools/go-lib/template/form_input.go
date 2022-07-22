package template

import liblanguage "github.com/feditools/go-lib/language"

type FormInputType string

const (
	// FormInputTypeCheckbox is a checkbox html input field.
	FormInputTypeCheckbox FormInputType = "checkbox"
	// FormInputTypeHidden is a hidden html input field.
	FormInputTypeHidden FormInputType = "hidden"
	// FormInputTypePassword is a password html input field.
	FormInputTypePassword FormInputType = "password"
	// FormInputTypeText is a text html input field.
	FormInputTypeText FormInputType = "text"
)

// FormInput is a templated form input.
type FormInput struct {
	ID           string
	Type         FormInputType
	Name         string
	Placeholder  string
	Label        *liblanguage.LocalizedString
	LabelClass   string
	Value        string
	WrappedClass string
	Disabled     bool
	Required     bool
	Checked      bool
	Validation   *FormValidation
}

// FormValidation is a validation response to a form input.
type FormValidation struct {
	Valid    bool
	Response string
}
