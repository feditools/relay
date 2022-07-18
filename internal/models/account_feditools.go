package models

import (
	"github.com/feditools/go-lib/fedihelper"
	"github.com/sirupsen/logrus"
	"time"
)

// GetActorURI returns the account's actor uri.
func (a *Account) GetActorURI() (actorURI string) {
	return a.ActorURI
}

// GetDisplayName returns the account's display name.
func (a *Account) GetDisplayName() (displayName string) {
	return a.DisplayName
}

// GetID returns the account's database id.
func (a *Account) GetID() (accountID int64) {
	return a.ID
}

// GetInstance returns the instance of the account.
func (a *Account) GetInstance() (instance fedihelper.Instance) {
	return a.Instance
}

// GetLastFinger returns the time of the last finger.
func (a *Account) GetLastFinger() (lastFinger time.Time) {
	return a.LastFinger
}

// GetUsername returns the account's username.
func (a *Account) GetUsername() (username string) {
	return a.ActorURI
}

// SetActorURI sets the account's actor uri.
func (a *Account) SetActorURI(actorURI string) {
	a.ActorURI = actorURI
}

// SetDisplayName sets the account's display name.
func (a *Account) SetDisplayName(displayName string) {
	a.DisplayName = displayName
}

// SetInstance sets the instance of the account.
func (a *Account) SetInstance(instanceI fedihelper.Instance) {
	l := logger.WithFields(logrus.Fields{
		"struct": "Account",
		"func":   "SetInstance",
	})

	instance, ok := instanceI.(*Instance)
	if !ok {
		l.Warnf("instance not type *FediInstance")

		return
	}

	a.InstanceID = instance.ID
	a.Instance = instance
}

// SetLastFinger sets the time of the last finger.
func (a *Account) SetLastFinger(lastFinger time.Time) {
	a.LastFinger = lastFinger
}

// SetUsername sets the account's username.
func (a *Account) SetUsername(username string) {
	a.Username = username
}
