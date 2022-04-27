package testdata

import (
	"github.com/feditools/relay/internal/models"
)

// TestInstances contains a set of test instances
var TestInstances = []*models.Instance{
	{
		ID:       1,
		Domain:   "example1.com",
		InboxIRI: "https://example1.com/inbox",
	},
	{
		ID:       2,
		Domain:   "example2.com",
		InboxIRI: "https://example3.com/inbox",
	},
	{
		ID:       3,
		Domain:   "example3.com",
		InboxIRI: "https://example3.com/inbox",
	},
	{
		ID:       4,
		Domain:   "example4.com",
		InboxIRI: "https://example4.com/inbox",
	},
	{
		ID:       5,
		Domain:   "example5.com",
		InboxIRI: "https://example5.com/inbox",
	},
}
