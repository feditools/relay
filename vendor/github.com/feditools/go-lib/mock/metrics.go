package mock

import (
	"time"

	"github.com/feditools/go-lib/metrics"
)

// GRPCRequest is a new database query metric measurer.
type GRPCRequest struct {
	start time.Time
}

// Done is called when the grpc request is complete.
func (g *GRPCRequest) Done(_ int) time.Duration {
	return time.Since(g.start)
}

// HTTPRequest is a new database query metric measurer.
type HTTPRequest struct {
	start time.Time
}

// Done is called when the http request is complete.
func (h *HTTPRequest) Done(_ int) time.Duration {
	return time.Since(h.start)
}

// DBQuery is a new database query metric measurer.
type DBQuery struct {
	start time.Time
}

// Done is called when the db query is complete.
func (d *DBQuery) Done(_ bool) time.Duration {
	return time.Since(d.start)
}

// DBCacheQuery is a new database cache query metric measurer.
type DBCacheQuery struct {
	start time.Time
}

// Done is called when the db cache query is complete.
func (d *DBCacheQuery) Done(_, _ bool) time.Duration {
	return time.Since(d.start)
}

// MetricsCollector is a mock metrics collection.
type MetricsCollector struct{}

// Close does nothing.
func (MetricsCollector) Close() error {
	return nil
}

// NewGRPCRequest creates a new grpc metrics collector.
func (MetricsCollector) NewGRPCRequest(_ string) metrics.GRPCRequest {
	return &GRPCRequest{
		start: time.Now(),
	}
}

// NewHTTPRequest creates a new http metrics collector.
func (MetricsCollector) NewHTTPRequest(_, _ string) metrics.HTTPRequest {
	return &HTTPRequest{
		start: time.Now(),
	}
}

// NewDBQuery creates a new db query metrics collector.
func (MetricsCollector) NewDBQuery(_ string) metrics.DBQuery {
	return &DBQuery{
		start: time.Now(),
	}
}

// NewDBCacheQuery creates a new db cache query metrics collector.
func (MetricsCollector) NewDBCacheQuery(_ string) metrics.DBCacheQuery {
	return &DBCacheQuery{
		start: time.Now(),
	}
}

// NewMetricsCollector creates a new mock metrics collector.
func NewMetricsCollector() (metrics.Collector, error) {
	return &MetricsCollector{}, nil
}
