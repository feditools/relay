package metrics

import "time"

// Collector collects metrics from the feditools.
type Collector interface {
	Close() error

	NewDBQuery(name string) DBQuery
	NewDBCacheQuery(name string) DBCacheQuery
	NewGRPCRequest(method string) GRPCRequest
	NewHTTPRequest(method, path string) HTTPRequest
}

// DBQuery is a new database query metric measurer.
type DBQuery interface {
	Done(isError bool) time.Duration
}

// DBCacheQuery is a new database cache query metric measurer.
type DBCacheQuery interface {
	Done(hit bool, isError bool) time.Duration
}

// GRPCRequest is a new grpc request metric measurer.
type GRPCRequest interface {
	Done(code int) time.Duration
}

// HTTPRequest is a new http request metric measurer.
type HTTPRequest interface {
	Done(status int) time.Duration
}
