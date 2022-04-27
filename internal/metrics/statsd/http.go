package statsd

import (
	"github.com/cactus/go-statsd-client/v5/statsd"
	"github.com/feditools/relay/internal/metrics"
	"strconv"
	"time"
)

// HTTPRequestTiming send a metrics relating to a http request
func (m *Module) HTTPRequestTiming(t time.Duration, status int, method, path string) {
	err := m.s.TimingDuration(
		metrics.StatHTTPRequest,
		t,
		m.rate,
		statsd.Tag{"status", strconv.Itoa(status)},
		statsd.Tag{"method", method},
		statsd.Tag{"path", path},
	)
	if err != nil {
		logger.WithField("func", "HTTPRequestTiming").Warn(err.Error())
	}
}
