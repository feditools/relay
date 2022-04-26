package action

import "context"

// Action defines an action that can be run
type Action func(context.Context) error
