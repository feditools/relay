package fedihelper

import (
	"context"
	"io"
	"net/http"
)

type HTTP interface {
	Do(req *http.Request) (resp *http.Response, err error)
	Get(ctx context.Context, url string) (resp *http.Response, err error)
	NewRequest(ctx context.Context, method, url string, body io.Reader) (req *http.Request, err error)
	Transport() (transport http.RoundTripper)
}
