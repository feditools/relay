package runner

import (
	"context"
	"github.com/feditools/go-lib/fedihelper"
)

type Runner interface {
	EnqueueInboxActivity(ctx context.Context, instanceID int64, actorIRI string, activity fedihelper.Activity) (err error)
	EnqueueDeliverActivity(ctx context.Context, instanceID int64, activity fedihelper.Activity) (err error)
}
