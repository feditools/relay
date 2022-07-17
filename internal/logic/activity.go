package logic

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/feditools/relay/internal/models"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"github.com/tyrm/go-util/mimetype"
	"net/url"
)

func (l *Logic) DeliverActivity(ctx context.Context, instanceID int64, activity models.Activity) error {
	log := logger.WithFields(logrus.Fields{
		"func": "DeliverActivity",
	})

	// get instance
	instance, err := l.db.ReadInstanceByID(ctx, instanceID)
	if err != nil {
		log.Errorf("db read: %s", err.Error())

		return fmt.Errorf("db read: %s", err.Error())
	}

	// send activity
	body, err := json.Marshal(activity)
	if err != nil {
		log.Errorf("can't marshal response: %s", err.Error())

		return fmt.Errorf("can't marshal response: %s", err.Error())
	}

	inboxIRI, err := url.Parse(instance.InboxIRI)
	if err != nil {
		log.Errorf("can't parse actor iri: %s", err.Error())

		return fmt.Errorf("can't parse actor iri: %s", err.Error())
	}

	log.Debugf("sending activity: %s", string(body))
	resp, err := l.transport.InstancePost(ctx, inboxIRI, body, mimetype.ApplicationActivityJSON, mimetype.ApplicationActivityJSON)
	if err != nil {
		log.Errorf("can't post to instance: %s\n%s", err.Error(), resp)

		return fmt.Errorf("can't post to instance: %s\n%s", err.Error(), resp)
	}

	return nil
}

func (l *Logic) ProcessActivity(ctx context.Context, instanceID int64, activity models.Activity) error {
	log := logger.WithFields(logrus.Fields{
		"func": "ProcessActivity",
	})

	actType, ok := activity["type"]
	if !ok {
		log.Debugf("activity missing type")

		return errors.New("activity missing type")
	}

	switch actType {
	case models.TypeAnnounce, models.TypeCreate:
		return l.doRelay(ctx, instanceID, activity)
	case models.TypeDelete, models.TypeUpdate:
		return l.doForward(ctx, instanceID, activity)
	case models.TypeFollow:
		return l.doFollow(ctx, instanceID, activity)
	case models.TypeUndo:
		return l.doUndo(ctx, instanceID, activity)
	default:
		log.Debugf("unhandled activity type: %s", actType)

		return fmt.Errorf("unhandled activity type: %s", actType)
	}
}

func (l *Logic) doFollow(ctx context.Context, instanceID int64, activity models.Activity) error {
	log := logger.WithFields(logrus.Fields{
		"func": "doFollow",
	})
	log.Trace("doFollow called")

	// get id
	idi, ok := activity["id"]
	if !ok {
		log.Debugf("activity missing id")

		return errors.New("activity missing id")
	}
	id, ok := idi.(string)
	if !ok {
		log.Debugf("activity id is not string")

		return errors.New("activity id is not string")
	}

	// set followed
	instance, err := l.db.ReadInstanceByID(ctx, instanceID)
	if err != nil {
		log.Errorf("db read: %s", err.Error())

		return fmt.Errorf("db read: %s", err.Error())
	}
	instance.Followed = true
	err = l.db.UpdateInstance(ctx, instance)
	if err != nil {
		log.Errorf("db update: %s", err.Error())

		return fmt.Errorf("db update: %s", err.Error())
	}

	// send accept
	outgoingActivity := genActivityAccept(l.domain, instance.ActorIRI, id)
	body, err := json.Marshal(outgoingActivity)
	if err != nil {
		log.Errorf("can't marshal response: %s", err.Error())

		return fmt.Errorf("can't marshal response: %s", err.Error())
	}

	inboxIRI, err := url.Parse(instance.InboxIRI)
	if err != nil {
		log.Errorf("can't parse actor iri: %s", err.Error())

		return fmt.Errorf("can't parse actor iri: %s", err.Error())
	}

	log.Debugf("sending activity: %s", string(body))
	resp, err := l.transport.InstancePost(ctx, inboxIRI, body, mimetype.ApplicationActivityJSON, mimetype.ApplicationActivityJSON)
	if err != nil {
		log.Errorf("can't post to instance: %s\n%s", err.Error(), resp)

		return fmt.Errorf("can't post to instance: %s\n%s", err.Error(), resp)
	}

	return nil
}

func (l *Logic) doForward(ctx context.Context, instanceID int64, activity models.Activity) error {
	log := logger.WithFields(logrus.Fields{
		"func": "doForward",
	})
	log.Trace("doForward called")

	// check if we've already forwarded
	objectID, err := activity.ObjectID()
	if err != nil {
		log.Warnf("object id: %s", err.Error())

		return fmt.Errorf("object id: %s", err.Error())
	}
	_, ok := l.cacheActivity.Get(objectID)
	if ok {
		log.Infof("already forwarded message: %v", objectID)

		return nil
	}
	objectIDURI, err := url.Parse(objectID)
	if err != nil {
		log.Errorf("parsing object id uri: %s", err.Error())

		return fmt.Errorf("parsing object id uri: %s", err.Error())
	}

	// get instance
	signingInstance, err := l.db.ReadInstanceByID(ctx, instanceID)
	if err != nil {
		log.Errorf("db read: %s", err.Error())

		return fmt.Errorf("db read: %s", err.Error())
	}

	// forward activity
	log.Debugf("forwarding messagge from %s", signingInstance.ActorIRI)
	log.Tracef("forwarding activity: %#v", activity)

	// read from db
	followedInstances, err := l.db.ReadInstancesWhereFollowing(ctx)
	if err != nil {
		log.Errorf("db read: %s", err.Error())

		return fmt.Errorf("db read: %s", err.Error())
	}
	log.Debugf("got %d followed instances", len(followedInstances))

	// enqueue deliveries
	for _, instance := range followedInstances {
		actorIRI, err := url.Parse(instance.ActorIRI)
		if err != nil {
			log.Errorf("parsing object id uri: %s", err.Error())

			return fmt.Errorf("parsing object id uri: %s", err.Error())
		}

		if instance.ActorIRI != signingInstance.ActorIRI && actorIRI.Host != objectIDURI.Host {
			err = l.runner.EnqueueDeliverActivity(ctx, instance.ID, activity)
			if err != nil {
				log.Errorf("enqueueing delivery: %s", err.Error())
			}
		}
	}
	_ = l.cacheActivity.Add(objectID, true)

	return nil
}

func (l *Logic) doRelay(ctx context.Context, instanceID int64, activity models.Activity) error {
	log := logger.WithFields(logrus.Fields{
		"func": "doRelay",
	})
	log.Trace("doRelay called")

	return nil
}

func (l *Logic) doUndo(ctx context.Context, instanceID int64, activity models.Activity) error {
	log := logger.WithFields(logrus.Fields{
		"func": "doUndo",
	})
	log.Trace("doUndo called")

	aType, err := activity.ObjectType()
	if err != nil {
		return err
	}

	switch aType {
	case models.TypeAnnounce:
		return l.doForward(ctx, instanceID, activity)
	case models.TypeFollow:
		// get instance
		instance, err := l.db.ReadInstanceByID(ctx, instanceID)
		if err != nil {
			log.Errorf("db read: %s", err.Error())

			return fmt.Errorf("db read: %s", err.Error())
		}

		// unset followed
		instance.Followed = false
		err = l.db.UpdateInstance(ctx, instance)
		if err != nil {
			log.Errorf("db update: %s", err.Error())

			return fmt.Errorf("db update: %s", err.Error())
		}

		// send accept
		outgoingActivity := genActivityUndo(l.domain, instance.ActorIRI)
		body, err := json.Marshal(outgoingActivity)
		if err != nil {
			log.Errorf("can't marshal response: %s", err.Error())

			return fmt.Errorf("can't marshal response: %s", err.Error())
		}

		inboxIRI, err := url.Parse(instance.InboxIRI)
		if err != nil {
			log.Errorf("can't parse actor iri: %s", err.Error())

			return fmt.Errorf("can't parse actor iri: %s", err.Error())
		}

		log.Debugf("sending activity: %s", string(body))
		resp, err := l.transport.InstancePost(ctx, inboxIRI, body, mimetype.ApplicationActivityJSON, mimetype.ApplicationActivityJSON)
		if err != nil {
			log.Errorf("can't post to instance: %s\n%s", err.Error(), resp)

			return fmt.Errorf("can't post to instance: %s\n%s", err.Error(), resp)
		}

		return nil
	default:
		log.Debugf("dropping activity object type: %s", aType)

		// drop activity
		return nil
	}
}

func genActivityID(domain string) string {
	return fmt.Sprintf("https://%s/activities/%s", domain, uuid.New().String())
}

func genActivityAccept(domain, to, activityID string) map[string]interface{} {
	return map[string]interface{}{
		"@context": models.ContextActivityStreams,
		"type":     models.TypeAccept,
		"to":       []string{to},
		"actor":    genActorSelf(domain),
		"object": map[string]interface{}{
			"type":   models.TypeFollow,
			"id":     activityID,
			"object": genActorSelf(domain),
			"actor":  to,
		},
		"id": genActivityID(domain),
	}
}

func genActivityUndo(domain, to string) map[string]interface{} {
	return map[string]interface{}{
		"@context": models.ContextActivityStreams,
		"type":     models.TypeUndo,
		"to":       []string{to},
		"actor":    genActorSelf(domain),
		"object": map[string]interface{}{
			"type":   models.TypeFollow,
			"id":     genActivityID(domain),
			"object": genActorSelf(domain),
			"actor":  genActorSelf(domain),
		},
		"id": genActivityID(domain),
	}
}
