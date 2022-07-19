package fedihelper

import (
	"github.com/go-fed/activity/pub"
	"net/http"
)

type HttpClient interface {
	pub.HttpClient
	Transport() (transport http.RoundTripper)
}
