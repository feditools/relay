package statsd

import (
	"fmt"
	"github.com/cactus/go-statsd-client/v5/statsd"
	"github.com/feditools/relay/internal/metrics"
	"time"
)

// DBQuery send a metrics relating to a database query
func (m *Module) DBQuery(t time.Duration, name string, isError bool) {
	l := logger.WithField("func", "DBQuery")

	err := m.s.TimingDuration(
		metrics.StatDBQueryTiming,
		t,
		m.rate,
		statsd.Tag{"name", name},
		statsd.Tag{"error", fmt.Sprintf("%v", isError)},
	)
	if err != nil {
		l.WithField("type", "timing").Warn(err.Error())
	}

	err = m.s.Inc(
		metrics.StatDBQueryCount,
		1,
		m.rate,
		statsd.Tag{"name", name},
		statsd.Tag{"error", fmt.Sprintf("%v", isError)},
	)
	if err != nil {
		l.WithField("type", "count").Warn(err.Error())
	}
}
