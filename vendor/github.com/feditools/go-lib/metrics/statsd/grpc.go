package statsd

import (
	"strconv"
	"time"

	"github.com/cactus/go-statsd-client/v5/statsd"
	"github.com/feditools/go-lib/metrics"
)

// GRPCRequest is a grpc request metric measurer.
type GRPCRequest struct {
	s    statsd.Statter
	rate float32

	method string
	start  time.Time
}

// NewGRPCRequest creates a new grpc request metrics collector.
func (m *Module) NewGRPCRequest(method string) metrics.GRPCRequest {
	return &GRPCRequest{
		s:    m.s,
		rate: m.rate,

		method: method,
		start:  time.Now(),
	}
}

// Done is called when the grpc request is complete.
func (g *GRPCRequest) Done(code int) time.Duration {
	l := logger.WithField("type", "GRPCRequest").WithField("func", "Done")

	t := time.Since(g.start)

	err := g.s.TimingDuration(
		metrics.StatGRPCRequestTiming,
		t,
		g.rate,
		statsd.Tag{metrics.TagMethod, g.method},
		statsd.Tag{metrics.TagCode, strconv.Itoa(code)},
	)
	if err != nil {
		l.WithField("kind", "timing").Warn(err.Error())
	}

	err = g.s.Inc(
		metrics.StatGRPCRequestCount,
		1,
		g.rate,
		statsd.Tag{metrics.TagMethod, g.method},
		statsd.Tag{metrics.TagCode, strconv.Itoa(code)},
	)
	if err != nil {
		l.WithField("kind", "count").Warn(err.Error())
	}

	return t
}
