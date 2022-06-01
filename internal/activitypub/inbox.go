package activitypub

import (
	"encoding/json"
	nethttp "net/http"
	"net/url"
)

func (m *Module) inboxPostHandler(w nethttp.ResponseWriter, r *nethttp.Request) {
	l := logger.WithField("func", "inboxPostHandler")

	// parse activity
	var activity map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&activity)
	if err != nil {
		l.Errorf("decoding activity: %+v", err)
		nethttp.Error(w, nethttp.StatusText(nethttp.StatusBadRequest), nethttp.StatusBadRequest)
		return
	}

	// parse actor uri
	actorI, ok := activity["actor"]
	if !ok {
		l.Debugf("activity missing actor: %+v", activity)
		nethttp.Error(w, nethttp.StatusText(nethttp.StatusBadRequest), nethttp.StatusBadRequest)
		return
	}
	actor, ok := actorI.(string)
	if !ok {
		l.Debugf("activity actor isn't string: %+v", activity)
		nethttp.Error(w, nethttp.StatusText(nethttp.StatusBadRequest), nethttp.StatusBadRequest)
		return
	}
	actorURI, err := url.Parse(actor)
	if err != nil {
		l.Errorf("can't parts actor uri from %s", actor)
		nethttp.Error(w, nethttp.StatusText(nethttp.StatusBadRequest), nethttp.StatusBadRequest)
		return
	}

	// check request validation
	validated, instance := m.validateRequest(r, actorURI)
	if !validated {
		l.Debugf("validation failed for actor: %s", actor)
		nethttp.Error(w, nethttp.StatusText(nethttp.StatusUnauthorized), nethttp.StatusUnauthorized)
		return
	}

	// get activity type
	activityTypeI, ok := activity["type"]
	if !ok {
		l.Debugf("activity missing type: %+v", activity)
		nethttp.Error(w, nethttp.StatusText(nethttp.StatusBadRequest), nethttp.StatusBadRequest)
		return
	}
	activityType, ok := activityTypeI.(string)
	if !ok {
		l.Debugf("activity type isn't string: %+v", activity)
		nethttp.Error(w, nethttp.StatusText(nethttp.StatusBadRequest), nethttp.StatusBadRequest)
		return
	}

	// drop non-Follow activities from instances that don't follow our relay
	if activityType != TypeFollow && !instance.Followed {
		nethttp.Error(w, nethttp.StatusText(nethttp.StatusUnauthorized), nethttp.StatusUnauthorized)
		return
	}

	l.Debugf("headers: %+v", r.Header)
	l.Debugf("body: %s", activity)
	w.WriteHeader(nethttp.StatusAccepted)
}
