package fedihelper

import (
	"context"
	"net/url"
)

type Instance interface {
	GetActorURI() (actorURI string)
	GetDomain() (domain string)
	GetID() (id int64)
	GetServerHostname() (hostname string)
	GetSoftware() (software string)

	SetActorURI(actorURI string)
	SetDomain(domain string)
	SetInboxURI(inboxURI string)
	SetServerHostname(hostname string)
	SetSoftware(software string)
}

// GenerateFediInstanceFromDomain created a Instance object by querying the apis of the federated instance.
func (f *FediHelper) GenerateFediInstanceFromDomain(ctx context.Context, domain string, instance Instance) error {
	l := logger.WithField("func", "GenerateFediInstanceFromDomain")

	// get host meta
	hostMeta, err := f.FetchHostMeta(ctx, domain)
	if err != nil {
		l.Errorf("get host meta: %s", err.Error())

		return err
	}
	hostMetaURIString, err := f.WebfingerURIFromHostMeta(hostMeta)
	if err != nil {
		l.Errorf("get webfinger uri: %s", err.Error())

		return err
	}
	hostMetaURI, err := url.Parse(hostMetaURIString)
	if err != nil {
		l.Errorf("parsing host meta uri: %s", err.Error())

		return err
	}

	// get nodeinfo endpoints from well-known location
	wkni, err := f.GetWellknownNodeInfo(ctx, hostMetaURI.Host)
	if err != nil {
		l.Errorf("get nodeinfo: %s", err.Error())

		return err
	}

	// check for nodeinfo 2.0 schema
	nodeinfoURI, err := findNodeInfo20URI(wkni)
	if err != nil {
		return err
	}
	if nodeinfoURI == nil {
		return NewError("missing nodeinfo 2.0 uri")
	}

	// get nodeinfo from
	nodeinfo, err := f.GetNodeInfo20(ctx, hostMetaURI.Host, nodeinfoURI)
	if err != nil {
		fhErr := NewErrorf("get nodeinfo 2.0: %s", err.Error())
		l.Error(fhErr.Error())

		return fhErr
	}

	// get actor uri
	webfinger, err := f.FetchWellknownWebFinger(ctx, hostMetaURI.Host, domain, domain)
	if err != nil {
		fhErr := NewErrorf("get wellknown webfinger: %s", err.Error())
		l.Error(fhErr.Error())

		return fhErr
	}
	actorURI, err := webfinger.ActorURI()
	if err != nil {
		fhErr := NewErrorf("find actor url: %s", err.Error())
		l.Error(fhErr.Error())

		return fhErr
	}
	if actorURI == nil {
		return NewError("missing actor uri")
	}

	instance.SetActorURI(actorURI.String())
	instance.SetDomain(domain)
	instance.SetServerHostname(hostMetaURI.Host)
	instance.SetSoftware(nodeinfo.Software.Name)

	return nil
}
