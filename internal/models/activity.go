package models

import (
	"errors"
	"github.com/sirupsen/logrus"
)

const (
	// ContextActivityStreams contains the context document for activity streams
	ContextActivityStreams = "https://www.w3.org/ns/activitystreams"

	// TypeAccept is the Accept activity Type
	TypeAccept = "Accept"
	// TypeAnnounce is the Announce activity Type
	TypeAnnounce = "Announce"
	// TypeCreate is the Create activity Type
	TypeCreate = "Create"
	// TypeDelete is the Delete activity Type
	TypeDelete = "Delete"
	// TypeFollow is the Follow activity Type
	TypeFollow = "Follow"
	// TypeUndo is the Undo activity Type
	TypeUndo = "Undo"
	// TypeUpdate is the Update activity Type
	TypeUpdate = "Update"
)

type Activity map[string]interface{}

func (a Activity) ObjectID() (string, error) {
	l := logger.WithFields(logrus.Fields{
		"func":  "ObjectID",
		"model": "Activity",
	})

	object, ok := a["object"]
	if !ok {
		return "", errors.New("activity is missing object")
	}

	switch o := object.(type) {
	case map[string]interface{}:
		idi, ok := o["id"]
		if !ok {
			return "", errors.New("object is missing id")
		}
		id, ok := idi.(string)
		if !ok {
			return "", errors.New("object id is not string")
		}

		return id, nil
	case string:
		return o, nil
	default:
		l.Warn("unknown object type")

		return "", errors.New("unknown object type")
	}
}

func (a Activity) ObjectType() (string, error) {
	objecti, ok := a["object"]
	if !ok {
		return "", errors.New("activity is missing object")
	}
	object, ok := objecti.(map[string]interface{})
	if !ok {
		return "", errors.New("activity object is wrong type")
	}

	aTypei, ok := object["type"]
	if !ok {
		return "", errors.New("activity object is missing type")
	}
	aType, ok := aTypei.(string)
	if !ok {
		return "", errors.New("activity object is wrong type")
	}

	return aType, nil
}
