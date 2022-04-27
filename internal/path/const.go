package path

const (
	// parts

	// PartActor is the noun used in a path for an actor
	PartActor = "actor"
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
	// PartWebFinger is the noun used in a path for web finger
	PartWebFinger = "webfinger"
	// PartWellKnown is the noun used in a well known path
	PartWellKnown = ".well-known"

	// activity pub

	// APActor is the path to the relay actor
	APActor = "/" + PartActor
	// APNodeInfo20 is the path to the node info 2.0 schema
	APNodeInfo20 = "/" + PartNodeinfo + "/2.0"
	// APWellKnownNodeInfo is the path to the well known node info endpoint
	APWellKnownNodeInfo = "/" + PartWellKnown + "/" + PartNodeinfo
	// APWellKnownWebFinger is the path to the well known web finger endpoint
	APWellKnownWebFinger = "/" + PartWellKnown + "/" + PartWebFinger
)
