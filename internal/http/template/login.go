package template

// LoginName is the name of the login template.
const LoginName = "login"

// Login contains the variables for the "login" template.
type Login struct {
	Common

	FormError   string
	FormAccount string
}
