package http

import "net/http"

// ResponseWriter is a ResponseWriter that keeps track of status and body size
type ResponseWriter struct {
	http.ResponseWriter
	status     int
	bodyLength int
}

// NewResponseWriter creates a new ResponseWriter
func NewResponseWriter(w http.ResponseWriter) *ResponseWriter {
	return &ResponseWriter{
		ResponseWriter: w,
		status:         200,
		bodyLength:     0,
	}
}

// BodyLength returns the response body length
func (r *ResponseWriter) BodyLength() int { return r.bodyLength }

// Status returns the status code of the response
func (r *ResponseWriter) Status() int { return r.status }

// Write to the response writer, also updating body length
func (r *ResponseWriter) Write(b []byte) (int, error) {
	n, err := r.ResponseWriter.Write(b)
	if err != nil {
		return 0, err
	}
	r.bodyLength += n
	return n, nil
}

// WriteHeader sets the status of the response
func (r *ResponseWriter) WriteHeader(status int) {
	r.ResponseWriter.WriteHeader(status)
	r.status = status
	return
}
