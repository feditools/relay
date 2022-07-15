package activitypub

import (
	"context"
	"encoding/json"
	"github.com/sirupsen/logrus"
	"github.com/tyrm/go-util/mimetype"
	"net/url"
	"sync"
)

type verifiedActivity struct {
	InstanceID int64
	Activity   map[string]interface{}
}

func (m *Module) worker(ctx context.Context, wid int, wg *sync.WaitGroup, inboxChan <-chan verifiedActivity, stop <-chan struct{}) {
	defer wg.Done()

	l := logger.WithFields(logrus.Fields{
		"func": "worker",
		"wid":  wid,
	})
	l.Debugf("starting worker")

	for {
		select {
		case vactivity, chok := <-inboxChan:
			if !chok {
				l.Debug("got channel close")

				return
			}

			actType, ok := vactivity.Activity["type"]
			if !ok {
				l.Debugf("activity missing type")

				continue
			}

			switch actType {
			case TypeFollow:
				m.doFollow(ctx, wid, &vactivity)
			default:
				l.Debugf("unhandled activity type: %s", actType)
			}
		case <-stop:
			l.Debug("got stop")

			return
		case <-ctx.Done():
			err := ctx.Err()
			l.Debugf("got context done: %s", err.Error())

			return
		}
	}
}

func (m *Module) doFollow(ctx context.Context, wid int, activity *verifiedActivity) {
	l := logger.WithFields(logrus.Fields{
		"func": "doFollow",
		"wid":  wid,
	})
	l.Trace("doFollow called")

	// get id
	idi, ok := activity.Activity["id"]
	if !ok {
		l.Debugf("activity missing id")

		return
	}
	id, ok := idi.(string)
	if !ok {
		l.Debugf("activity id is not string")

		return
	}

	// set followed
	instance, err := m.db.ReadInstanceByID(ctx, activity.InstanceID)
	if err != nil {
		l.Errorf("db read: %s", err.Error())

		return
	}
	instance.Followed = true
	err = m.db.UpdateInstance(ctx, instance)
	if err != nil {
		l.Errorf("db update: %s", err.Error())

		return
	}

	// send accept
	outgoingActivity := m.genActivityAccept(instance.ActorIRI, id)
	body, err := json.Marshal(outgoingActivity)
	if err != nil {
		l.Errorf("can't marshal response: %s", err.Error())

		return
	}

	inboxIRI, err := url.Parse(instance.InboxIRI)
	if err != nil {
		l.Errorf("can't parse actor iri: %s", err.Error())

		return
	}

	l.Debugf("sending activity: %s", string(body))
	resp, err := m.transport.InstancePost(ctx, inboxIRI, body, mimetype.ApplicationActivityJSON, mimetype.ApplicationActivityJSON)
	if err != nil {
		l.Errorf("can't post to instance: %s\n%s", err.Error(), resp)

		return
	}
}
