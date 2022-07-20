package testdata

import (
	"github.com/feditools/relay/internal/models"
)

// TestInstances contains a set of test instances
var TestInstances = []*models.Instance{
	{
		ID:             1,
		ServerHostname: "example1.com",
		InboxIRI:       "https://example1.com/inbox",
	},
	{
		ID:             2,
		ServerHostname: "example2.com",
		InboxIRI:       "https://example3.com/inbox",
	},
	{
		ID:             3,
		ServerHostname: "example3.com",
		InboxIRI:       "https://example3.com/inbox",
	},
	{
		ID:             4,
		ServerHostname: "example4.com",
		InboxIRI:       "https://example4.com/inbox",
	},
	{
		ID:             5,
		ServerHostname: "example5.com",
		InboxIRI:       "https://example5.com/inbox",
	},
}
