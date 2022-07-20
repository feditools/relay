package path

const (
	// public

	// App is the path for the web app.
	App = "/" + PartApp

	// AppCallbackOauth is the path for an oauth callback.
	AppCallbackOauth = App + AppSubCallbackOauth
	// AppSubCallbackOauth is the sub path for an oauth callback.
	AppSubCallbackOauth = "/" + PartCallback + "/" + PartOauth + "/" + VarInstance
	// AppPreCallbackOauth is the prefix path for an oauth callback.
	AppPreCallbackOauth = App + "/" + PartCallback + "/" + PartOauth + "/"
	// AppHome is the path for the home page.
	AppHome = App + AppSubHome
	// AppSubHome is the sub path for the home page.
	AppSubHome = "/"
	// AppLogin is the path for the login page.
	AppLogin = App + AppSubLogin
	// AppSubLogin is the sub path for the login page.
	AppSubLogin = "/" + PartLogin
	// AppLogout is the path for the logout page.
	AppLogout = App + AppSubLogout
	// AppSubLogout is the sub path for the logout page.
	AppSubLogout = "/" + PartLogout

	// admin

	// AppAdmin is the path for the admin page.
	AppAdmin = "/" + PartAdmin
	// AppAdminHome is the path for the home page.
	AppAdminHome = App + AppAdmin + AppAdminSubHome
	// AppAdminSubHome is the sub path for the home page.
	AppAdminSubHome = "/"
	// AppAdminInstances is the path for the admin instances page.
	AppAdminInstances = App + AppAdmin + AppAdminSubInstances
	// AppAdminSubInstances is the sub path for the admin instances page.
	AppAdminSubInstances = "/" + PartInstance
)
