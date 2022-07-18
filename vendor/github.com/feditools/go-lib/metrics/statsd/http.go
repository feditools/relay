package statsd

import (
	"strconv"
	"time"

	"github.com/cactus/go-statsd-client/v5/statsd"
	"github.com/feditools/go-lib/metrics"
)

// HTTPRequest is a new http request metric measurer.
type HTTPRequest struct {
	s    statsd.Statter
	rate float32

	method string
	path   string
	start  time.Time
}

// NewHTTPRequest creates a new db query metrics collector.
func (m *Module) NewHTTPRequest(method, path string) metrics.HTTPRequest {
	return &HTTPRequest{
		s:    m.s,
		rate: m.rate,

		method: method,
		path:   path,
		start:  time.Now(),
	}
}

// Done is called when the db query is complete.
func (h *HTTPRequest) Done(status int) time.Duration {
	l := logger.WithField("type", "HTTPRequest").WithField("func", "Done")

	t := time.Since(h.start)

	err := h.s.TimingDuration(
		metrics.StatHTTPRequestTiming,
		t,
		h.rate,
		statsd.Tag{metrics.TagStatus, strconv.Itoa(status)},
		statsd.Tag{metrics.TagMethod, h.method},
		statsd.Tag{metrics.TagPath, h.path},
	)
	if err != nil {
		l.WithField("kind", "timing").Warn(err.Error())
	}

	err = h.s.Inc(
		metrics.StatHTTPRequestCount,
		1,
		h.rate,
		statsd.Tag{metrics.TagStatus, strconv.Itoa(status)},
		statsd.Tag{metrics.TagMethod, h.method},
		statsd.Tag{metrics.TagPath, h.path},
	)
	if err != nil {
		l.WithField("kind", "count").Warn(err.Error())
	}

	return t
}
