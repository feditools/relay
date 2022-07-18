package runner

import (
	"context"
	"github.com/feditools/relay/internal/models"
)

type Runner interface {
	EnqueueInboxActivity(ctx context.Context, instanceID int64, actorIRI string, activity models.Activity) (err error)
	EnqueueDeliverActivity(ctx context.Context, instanceID int64, activity models.Activity) (err error)
}
