package metrics

import "time"

// Collector collects metrics and emits them
type Collector interface {
	Close() error
	DBQuery(t time.Duration, name string, error bool)
	HTTPRequestTiming(t time.Duration, status int, method, path string)
}
