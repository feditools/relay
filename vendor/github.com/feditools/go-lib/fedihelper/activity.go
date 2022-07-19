package fedihelper

import (
	"errors"
	"github.com/sirupsen/logrus"
)

type Activity map[string]interface{}

func (a Activity) ID() (string, error) {
	idi, ok := a["id"]
	if !ok {
		return "", errors.New("activity is missing id")
	}
	id, ok := idi.(string)
	if !ok {
		return "", errors.New("activity id is wrong type")
	}

	return id, nil
}

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

func (a Activity) ObjectType() (ActivityType, error) {
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

	return ActivityType(aType), nil
}

func (a Activity) Type() (ActivityType, error) {
	aTypei, ok := a["type"]
	if !ok {
		return "", errors.New("activity is missing type")
	}
	aType, ok := aTypei.(string)
	if !ok {
		return "", errors.New("activity type is wrong type")
	}

	return ActivityType(aType), nil
}
