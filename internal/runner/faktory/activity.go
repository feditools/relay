package faktory

import (
	"context"
	"fmt"
	faktory "github.com/contribsys/faktory/client"
	worker "github.com/contribsys/faktory_worker_go"
	"github.com/feditools/relay/internal/models"
	"github.com/sirupsen/logrus"
	"strconv"
)

func (r *Runner) EnqueueDeliverActivity(_ context.Context, instanceID int64, activity models.Activity) error {
	job := faktory.NewJob(JobDeliverActivity, strconv.FormatInt(instanceID, 10), activity)
	job.Queue = QueueDelivery

	client, err := r.manager.Pool.Get()
	if err != nil {
		return err
	}
	return client.Push(job)
}

func (r *Runner) deliverActivity(ctx context.Context, args ...interface{}) error {
	help := worker.HelperFor(ctx)

	l := logger.WithFields(logrus.Fields{
		"func": "deliverActivity",
		"jid":  help.Jid(),
	})

	if len(args) != 2 {
		l.Errorf("wrong number of arguments, got: %d, want: %d", len(args), 2)
	}

	// cast arguments
	instanceIDStr, ok := args[0].(string)
	if !ok {
		l.Errorf("argument 0 is not an string")

		return fmt.Errorf("argument 0 is not an int")
	}
	instanceID, err := strconv.ParseInt(instanceIDStr, 10, 64)
	if err != nil {
		l.Errorf("cant parse int from argument 0: %s", err.Error())

		return fmt.Errorf("cant parse int from argument 0: %s", err.Error())

	}
	activity, ok := args[1].(map[string]interface{})
	if !ok {
		l.Errorf("argument 1 is not an activity")

		return fmt.Errorf("argument 1 is not an activity")
	}

	return r.logic.DeliverActivity(ctx, instanceID, activity)
}

func (r *Runner) EnqueueInboxActivity(_ context.Context, instanceID int64, activity models.Activity) error {
	job := faktory.NewJob(JobInboxActivity, strconv.FormatInt(instanceID, 10), activity)
	job.Queue = QueueDefault

	client, err := r.manager.Pool.Get()
	if err != nil {
		return err
	}
	return client.Push(job)
}

func (r *Runner) inboxActivity(ctx context.Context, args ...interface{}) error {
	help := worker.HelperFor(ctx)

	l := logger.WithFields(logrus.Fields{
		"func": "inboxActivity",
		"jid":  help.Jid(),
	})

	if len(args) != 2 {
		l.Errorf("wrong number of arguments, got: %d, want: %d", len(args), 2)
	}

	// cast arguments
	instanceIDStr, ok := args[0].(string)
	if !ok {
		l.Errorf("argument 0 is not an string")

		return fmt.Errorf("argument 0 is not an int")
	}
	instanceID, err := strconv.ParseInt(instanceIDStr, 10, 64)
	if err != nil {
		l.Errorf("cant parse int from argument 0: %s", err.Error())

		return fmt.Errorf("cant parse int from argument 0: %s", err.Error())

	}
	activity, ok := args[1].(map[string]interface{})
	if !ok {
		l.Errorf("argument 1 is not an activity")

		return fmt.Errorf("argument 1 is not an activity")
	}

	// process activity
	return r.logic.ProcessActivity(ctx, instanceID, activity)
}
