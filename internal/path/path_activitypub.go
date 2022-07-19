package path

const (
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
)
