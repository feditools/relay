package template

import "github.com/feditools/relay/internal/models"

// HomeName is the name of the home template.
const HomeName = "home"

// Home contains the variables for the home template.
type Home struct {
	Common

	ActorHref          string
	InboxHref          string
	FollowingInstances []*models.Instance
	Blocks             []*models.Block
}
