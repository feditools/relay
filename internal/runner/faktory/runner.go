package faktory

import (
	"context"
	worker "github.com/contribsys/faktory_worker_go"
	"github.com/feditools/relay/internal/config"
	"github.com/feditools/relay/internal/logic/logic1"
	"github.com/spf13/viper"
)

type Runner struct {
	logic   *logic1.Logic
	manager *worker.Manager
}

// New created a new logic module
func New(l *logic1.Logic) (*Runner, error) {
	newRunner := &Runner{
		logic: l,
	}

	mgr := worker.NewManager()
	mgr.Concurrency = viper.GetInt(config.Keys.RunnerConcurrency)
	mgr.ProcessWeightedPriorityQueues(map[string]int{QueueDefault: 2, QueueDelivery: 1})

	mgr.Register(JobDeliverActivity, newRunner.deliverActivity)
	mgr.Register(JobInboxActivity, newRunner.inboxActivity)

	newRunner.manager = mgr

	return newRunner, nil
}

func (r *Runner) Start(ctx context.Context) {
	l := logger.WithField("func", "Start")

	go func() {
		err := r.manager.RunWithContext(ctx)
		if err != nil {
			l.Errorf("run error: %s", err.Error())
		}
	}()
}
