package activitypub

import (
	"encoding/json"
	"github.com/feditools/relay/internal/models"
	nethttp "net/http"
	"net/url"
)

func (m *Module) inboxPostHandler(w nethttp.ResponseWriter, r *nethttp.Request) {
	l := logger.WithField("func", "inboxPostHandler")

	// parse activity
	var activity models.Activity
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
	actorStr, ok := actorI.(string)
	if !ok {
		l.Debugf("activity actor isn't string: %+v", activity)
		nethttp.Error(w, nethttp.StatusText(nethttp.StatusBadRequest), nethttp.StatusBadRequest)
		return
	}
	actorURI, err := url.Parse(actorStr)
	if err != nil {
		l.Errorf("can't parts actor uri from %s: %s", actorStr, err.Error())
		nethttp.Error(w, nethttp.StatusText(nethttp.StatusBadRequest), nethttp.StatusBadRequest)
		return
	}

	// check request validation
	l.Tracef("validating actor: %s", actorStr)
	validated, actor := m.logic.ValidateRequest(r, actorURI)
	if !validated {
		l.Debugf("validation failed for actor: %s", actor)
		nethttp.Error(w, nethttp.StatusText(nethttp.StatusUnauthorized), nethttp.StatusUnauthorized)
		return
	}

	// get instance of actor
	instance, err := m.logic.GetInstance(r.Context(), actorURI.Host)
	if err != nil {
		l.Errorf("can't get instance %s: %s", actorStr, err.Error())
		nethttp.Error(w, nethttp.StatusText(nethttp.StatusInternalServerError), nethttp.StatusInternalServerError)
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
	if activityType != models.TypeFollow && !instance.Followed {
		l.Debugf("got non follow from an unfollowed instance: %s", activityType)
		nethttp.Error(w, nethttp.StatusText(nethttp.StatusUnauthorized), nethttp.StatusUnauthorized)
		return
	}

	l.Debugf("headers: %+v", r.Header)
	l.Debugf("body: %s", activity)

	// enqueue activity
	err = m.runner.EnqueueInboxActivity(r.Context(), instance.ID, activity)
	if err != nil {
		nethttp.Error(w, nethttp.StatusText(nethttp.StatusInternalServerError), nethttp.StatusInternalServerError)
		return
	}

	w.WriteHeader(nethttp.StatusAccepted)
}
