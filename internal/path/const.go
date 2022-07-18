package path

const (
	// files.

	// FileDefaultCSS is the css document applies to all pages.
	FileDefaultCSS = StaticCSS + "/default.min.css"
	// FileErrorCSS is the css document applies to the error page.
	FileErrorCSS = StaticCSS + "/error.min.css"
	// FileLoginCSS is the css document applies to the login page.
	FileLoginCSS = StaticCSS + "/login.min.css"

	// parts

	// PartActor is the noun used in a path for an actor
	PartActor = "actor"
	// PartAdmin is used in a path for administrative tasks.
	PartAdmin = "admin"
	// PartFollowers is the noun used in a path for an actor's followers
	PartFollowers = "followers"
	// PartFollowing is the noun used in a path for actors following an actor
	PartFollowing = "following"
	// PartInbox is the noun used in a path for an actor's inbox
	PartInbox = "inbox"
	// PartNodeinfo is the noun used in a path for a node info object
	PartNodeinfo = "nodeinfo"
	// PartPublicKey is the noun used in a path for an actor's public key
	PartPublicKey = "main-key"
	// PartStatic is used in a path for static files.
	PartStatic = "static"
	// PartWebFinger is the noun used in a path for web finger
	PartWebFinger = "webfinger"
	// PartWellKnown is the noun used in a well known path
	PartWellKnown = ".well-known"

	// activity pub

	// APActor is the path to the relay actor
	APActor = "/" + PartActor
	// APInbox is the path to the relay inbox
	APInbox = "/" + PartInbox
	// APNodeInfo20 is the path to the node info 2.0 schema
	APNodeInfo20 = "/" + PartNodeinfo + "/2.0"
	// APWellKnownNodeInfo is the path to the well known node info endpoint
	APWellKnownNodeInfo = "/" + PartWellKnown + "/" + PartNodeinfo
	// APWellKnownWebFinger is the path to the well known web finger endpoint
	APWellKnownWebFinger = "/" + PartWellKnown + "/" + PartWebFinger

	// web app

	// Admin is the path for the admin page.
	Admin = "/" + PartAdmin
	// Home is the path for the home page.
	Home = "/"
	// Static is the path for static files.
	Static = "/" + PartStatic + "/"
	// StaticCSS is the path.
	StaticCSS = Static + "css"
)
