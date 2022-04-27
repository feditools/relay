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
	// PartPublicKey is the noun used in a path for an actor's public key
	PartPublicKey = "main-key"

	// activity pub

	// APActor is the path to the relay actor
	APActor = "/" + PartActor
)
