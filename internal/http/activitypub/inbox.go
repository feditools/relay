package activitypub

import (
	"encoding/json"
	"errors"
	"github.com/feditools/go-lib/fedihelper"
	"github.com/feditools/relay/internal/db"
	"github.com/feditools/relay/internal/models"
	nethttp "net/http"
	"net/url"
)

func (m *Module) inboxPostHandler(w nethttp.ResponseWriter, r *nethttp.Request) {
	l := logger.WithField("func", "inboxPostHandler")

	// parse activity
	var activity fedihelper.Activity
	err := json.NewDecoder(r.Body).Decode(&activity)
	if err != nil {
		l.Errorf("decoding activity: %+v", err)
		nethttp.Error(w, nethttp.StatusText(nethttp.StatusBadRequest), nethttp.StatusBadRequest)

		return
	}
	l.Tracef("got activity: %#v", activity)

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
	actorIRI, err := url.Parse(actorStr)
	if err != nil {
		l.Errorf("can't parts actor uri from %s: %s", actorStr, err.Error())
		nethttp.Error(w, nethttp.StatusText(nethttp.StatusBadRequest), nethttp.StatusBadRequest)

		return
	}

	// check request validation
	l.Tracef("validating actor: %s", actorStr)
	validated, actor := m.logic.ValidateRequest(r, actorIRI)
	if !validated {
		l.Debugf("validation failed for actor: %s", actor)
		nethttp.Error(w, nethttp.StatusText(nethttp.StatusUnauthorized), nethttp.StatusUnauthorized)

		return
	}

	var instance *models.Instance
	switch actor.Type {
	case fedihelper.TypeApplication:
		instance, err = m.logic.GetInstanceForActor(r.Context(), actorIRI)
		if err != nil {
			if errors.Is(err, db.ErrNoEntries) {
				nethttp.Error(w, nethttp.StatusText(nethttp.StatusUnauthorized), nethttp.StatusUnauthorized)

				return
			}
			l.Errorf("can't get instance %s: %s", actorStr, err.Error())
			nethttp.Error(w, nethttp.StatusText(nethttp.StatusInternalServerError), nethttp.StatusInternalServerError)

			return
		}
	case fedihelper.TypePerson:
		instance, err = m.logic.GetInstance(r.Context(), actorIRI.Host)
		if err != nil {
			if errors.Is(err, db.ErrNoEntries) {
				nethttp.Error(w, nethttp.StatusText(nethttp.StatusUnauthorized), nethttp.StatusUnauthorized)

				return
			}
			l.Errorf("can't get instance %s: %s", actorStr, err.Error())
			nethttp.Error(w, nethttp.StatusText(nethttp.StatusInternalServerError), nethttp.StatusInternalServerError)

			return
		}
	default:
		l.Errorf("unknown actor type: %s", actor.Type)
		nethttp.Error(w, nethttp.StatusText(nethttp.StatusBadRequest), nethttp.StatusBadRequest)

		return
	}

	// get activity type
	activityType, err := activity.Type()
	if err != nil {
		l.Debugf("can't get type: %s", err.Error())
		nethttp.Error(w, nethttp.StatusText(nethttp.StatusBadRequest), nethttp.StatusBadRequest)

		return
	}

	// drop non-Follow activities from instances that don't follow our relay
	if activityType != fedihelper.TypeFollow && !instance.Followed {
		l.Debugf("got non follow from an unfollowed instance: %s", activityType)
		nethttp.Error(w, nethttp.StatusText(nethttp.StatusUnauthorized), nethttp.StatusUnauthorized)

		return
	}

	l.Debugf("headers: %+v", r.Header)
	l.Debugf("body: %s", activity)

	// enqueue activity
	err = m.runner.EnqueueInboxActivity(r.Context(), instance.ID, actorStr, activity)
	if err != nil {
		nethttp.Error(w, nethttp.StatusText(nethttp.StatusInternalServerError), nethttp.StatusInternalServerError)

		return
	}

	w.WriteHeader(nethttp.StatusAccepted)
}
