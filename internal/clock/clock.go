package clock

import (
	"github.com/go-fed/activity/pub"
	"time"
)

// Clock implements the Clock interface of go-fed
type Clock struct{}

// Now  returns the current time
func (c *Clock) Now() time.Time {
	return time.Now()
}

// NewClock returns a simple pub.Clock for use in federation interfaces.
func NewClock() pub.Clock {
	return &Clock{}
}
