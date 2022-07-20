package fedihelper

import (
	"time"

	"golang.org/x/sync/singleflight"
)

const actorCacheExp = 60 * time.Minute
const nodeInfoCacheExp = 60 * time.Minute

// New creates a new fedi module.
func New(k KV, t *Transport, clientName string, helpers []Helper) (*FediHelper, error) {
	newFedi := &FediHelper{
		http: t,
		kv:   k,

		helpers: map[SoftwareName]Helper{},

		appClientName:    clientName,
		actorCacheExp:    actorCacheExp,
		nodeinfoCacheExp: nodeInfoCacheExp,
	}

	// add helpers
	for _, h := range helpers {
		h.SetFedi(newFedi)
		newFedi.helpers[h.GetSoftware()] = h
	}

	return newFedi, nil
}

// FediHelper is a module for working with federated social instances.
type FediHelper struct {
	http *Transport
	kv   KV

	CreateAccountHandler CreateAccountHandler
	GetAccountHandler    GetAccountHandler
	NewAccountHandler    NewAccountHandler

	helpers map[SoftwareName]Helper

	appClientName    string
	actorCacheExp    time.Duration
	nodeinfoCacheExp time.Duration
	requestGroup     singleflight.Group
}

func (f *FediHelper) SetCreateAccountHandler(handler CreateAccountHandler) {
	f.CreateAccountHandler = handler
}

func (f *FediHelper) SetGetAccountHandler(handler GetAccountHandler) {
	f.GetAccountHandler = handler
}

func (f *FediHelper) SetNewAccountHandler(handler NewAccountHandler) {
	f.NewAccountHandler = handler
}
