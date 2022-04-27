package mock

import (
	"github.com/feditools/relay/internal/metrics"
	"time"
)

// MetricsCollector is a mock metrics collection
type MetricsCollector struct{}

// NewMetricsCollector creates a new mock metrics collector
func NewMetricsCollector() (metrics.Collector, error) {
	return &MetricsCollector{}, nil
}

// Close does nothing
func (m MetricsCollector) Close() error {
	return nil
}

// DBQuery does nothing
func (m MetricsCollector) DBQuery(t time.Duration, name string, error bool) {
	return
}

// HTTPRequestTiming does nothing
func (m MetricsCollector) HTTPRequestTiming(t time.Duration, status int, method, path string) {
	return
}
