package fedihelper

import (
	"time"

	"golang.org/x/sync/singleflight"
)

const nodeInfoCacheExp = 60 * time.Minute

// New creates a new fedi module.
func New(h HTTP, k KV, clientName string, helpers []Helper) (*FediHelper, error) {
	newFedi := &FediHelper{
		http: h,
		kv:   k,

		helpers: map[Software]Helper{},

		appClientName:    clientName,
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
	http HTTP
	kv   KV

	CreateAccountHandler  CreateAccountHandler
	CreateInstanceHandler CreateInstanceHandler
	GetAccountHandler     GetAccountHandler
	GetInstanceHandler    GetInstanceHandler
	GetTokenHandler       GetTokenHandler
	NewAccountHandler     NewAccountHandler
	NewInstanceHandler    NewInstanceHandler
	UpdateInstanceHandler UpdateInstanceHandler

	helpers map[Software]Helper

	appClientName    string
	nodeinfoCacheExp time.Duration
	requestGroup     singleflight.Group
}

func (f *FediHelper) HTTP() HTTP {
	return f.http
}

func (f *FediHelper) SetCreateAccountHandler(handler CreateAccountHandler) {
	f.CreateAccountHandler = handler
}

func (f *FediHelper) SetCreateInstanceHandler(handler CreateInstanceHandler) {
	f.CreateInstanceHandler = handler
}

func (f *FediHelper) SetGetAccountHandler(handler GetAccountHandler) {
	f.GetAccountHandler = handler
}

func (f *FediHelper) SetGetInstanceHandler(handler GetInstanceHandler) {
	f.GetInstanceHandler = handler
}

func (f *FediHelper) SetGetTokenHandler(handler GetTokenHandler) {
	f.GetTokenHandler = handler
}

func (f *FediHelper) SetNewAccountHandler(handler NewAccountHandler) {
	f.NewAccountHandler = handler
}

func (f *FediHelper) SetNewInstanceHandler(handler NewInstanceHandler) {
	f.NewInstanceHandler = handler
}

func (f *FediHelper) SetUpdateInstanceHandler(handler UpdateInstanceHandler) {
	f.UpdateInstanceHandler = handler
}
