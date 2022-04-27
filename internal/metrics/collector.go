package metrics

import "time"

// Collector collects metrics and emits them
type Collector interface {
	Close() error
	HTTPRequestTiming(t time.Duration, status int, method, path string)
}
