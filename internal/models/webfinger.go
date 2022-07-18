package models

import (
	"net/url"
)

// WebFinger represents a web finger response
type WebFinger struct {
	Aliases []string `json:"aliases,omitempty"`
	Links   []Link   `json:"links,omitempty"`
	Subject string   `json:"subject,omitempty"`
}

// ActorURI an actor uri.
func (w *WebFinger) ActorURI() (*url.URL, error) {
	var actorURIstr string
	for _, link := range w.Links {
		if link.Rel == "self" || link.Type == MimeAppActivityJSON {
			actorURIstr = link.Href

			break
		}
	}
	if actorURIstr == "" {
		return nil, NewError("missing actor uri")
	}

	actorURI, err := url.Parse(actorURIstr)
	if err != nil {
		return nil, NewErrorf("invalid actor uri: %s", err.Error())
	}

	return actorURI, err
}
