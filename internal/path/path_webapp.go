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
	// AppAdminBlockList is the path for the admin Blocks page.
	AppAdminBlockList = App + AppAdmin + AppAdminSubBlockList
	// AppAdminSubBlockList is the sub path for the admin Blocks page.
	AppAdminSubBlockList = "/" + PartBlock
	// AppAdminBlockView is the path for the admin Block view page.
	AppAdminBlockView = App + AppAdmin + AppAdminSubBlockView
	// AppAdminPreBlockView is the prefix path for for the admin Block view page.
	AppAdminPreBlockView = App + AppAdmin + "/" + PartBlock + "/"
	// AppAdminSubBlockView is the sub path for the admin Block view page.
	AppAdminSubBlockView = "/" + PartBlock + "/" + VarBlock
	// AppAdminHome is the path for the home page.
	AppAdminHome = App + AppAdmin + AppAdminSubHome
	// AppAdminSubHome is the sub path for the home page.
	AppAdminSubHome = "/"
	// AppAdminInstanceDeleteFTmpl is the sub path for the admin instance delete action.
	AppAdminInstanceDeleteFTmpl = "/" + PartInstance + "/" + VarInstance + "/" + PartDelete
	// AppAdminSubInstanceDelete is the sub path for the admin instance delete action.
	AppAdminSubInstanceDelete = "/" + PartInstance + "/" + VarInstance + "/" + PartDelete
	// AppAdminInstanceList is the path for the admin instances page.
	AppAdminInstanceList = App + AppAdmin + AppAdminSubInstanceList
	// AppAdminSubInstanceList is the sub path for the admin instances page.
	AppAdminSubInstanceList = "/" + PartInstance
	// AppAdminInstanceView is the path for the admin instance view page.
	AppAdminInstanceView = App + AppAdmin + AppAdminSubInstanceView
	// AppAdminPreInstanceView is the prefix path for for the admin instance view page.
	AppAdminPreInstanceView = App + AppAdmin + "/" + PartInstance + "/"
	// AppAdminSubInstanceView is the sub path for the admin instance view page.
	AppAdminSubInstanceView = "/" + PartInstance + "/" + VarInstance
)
