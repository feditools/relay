package path

const (
	// public

	// App is the path for the web app.
	App = "/" + PartApp
	// AppHome is the path for the home page.
	AppHome = App + AppSubHome
	// AppSubHome is the path for the home page.
	AppSubHome = "/"
	// AppLogin is the path for the login page.
	AppLogin = App + AppSubLogin
	// AppSubLogin is the path for the login page.
	AppSubLogin = "/" + PartLogin
	// AppLogout is the path for the logout page.
	AppLogout = App + AppSubLogout
	// AppSubLogout is the path for the logout page.
	AppSubLogout = "/" + PartLogout

	// admin

	// AppAdmin is the path for the admin page.
	AppAdmin = "/" + PartAdmin
	// AppAdminHome is the path for the home page.
	AppAdminHome = App + AppAdmin + AppAdminSubHome
	// AppAdminSubHome is the path for the home page.
	AppAdminSubHome = "/"
)
