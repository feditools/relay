package faktory

import (
	"context"
	"fmt"
	faktory "github.com/contribsys/faktory/client"
	worker "github.com/contribsys/faktory_worker_go"
	"github.com/sirupsen/logrus"
	"strconv"
)

func (r *Runner) EnqueueProcessBlock(ctx context.Context, blockID int64) error {
	job := faktory.NewJob(JobProcessBlock, strconv.FormatInt(blockID, 10))
	job.Queue = QueueDelivery

	client, err := r.manager.Pool.Get()
	if err != nil {
		return err
	}
	return client.Push(job)
}

func (r *Runner) processBlock(ctx context.Context, args ...interface{}) error {
	help := worker.HelperFor(ctx)

	l := logger.WithFields(logrus.Fields{
		"func": "processBlock",
		"jid":  help.Jid(),
	})

	if len(args) != 1 {
		l.Errorf("wrong number of arguments, got: %d, want: %d", len(args), 2)
	}

	// cast arguments
	blockIDStr, ok := args[0].(string)
	if !ok {
		l.Errorf("argument 0 is not an string")

		return fmt.Errorf("argument 0 is not an string")
	}
	blockID, err := strconv.ParseInt(blockIDStr, 10, 64)
	if err != nil {
		l.Errorf("cant parse int from argument 0: %s", err.Error())

		return fmt.Errorf("cant parse int from argument 0: %s", err.Error())
	}

	return r.logic.ProcessBlock(ctx, help.Jid(), blockID)
}
