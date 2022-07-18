package http

import "net/http"

// Transport adds the expected http User-Agent to any request
type Transport struct {
}

// RoundTrip executes the default http.Transport with expected http User-Agent
func (*Transport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Set("User-Agent", GetUserAgent())
	return http.DefaultTransport.RoundTrip(req)
}
