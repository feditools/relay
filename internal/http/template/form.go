package template

type FormInput struct {
	Class       string
	ID          string
	Name        string
	Placeholder string
	Type        string
	Value       string

	Validation *FormInputValidation
}

type FormInputValidation struct {
	Message string
	Valid   bool
}
