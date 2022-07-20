package models

func (i *Instance) GetActorURI() (actorURI string) {
	return i.ActorIRI
}

func (i *Instance) GetDomain() (domain string) {
	return i.AccountDomain
}

func (i *Instance) GetID() (instanceID int64) {
	return i.ID
}

func (i *Instance) GetServerHostname() (hostname string) {
	return i.ServerHostname
}

func (i *Instance) GetSoftware() (software string) {
	return i.Software
}

func (i *Instance) SetActorURI(actorURI string) {
	i.ActorIRI = actorURI
}

func (i *Instance) SetDomain(domain string) {
	i.AccountDomain = domain
}

func (i *Instance) SetInboxURI(inboxURI string) {
	i.InboxIRI = inboxURI
}

func (i *Instance) SetServerHostname(hostname string) {
	i.ServerHostname = hostname
}

func (i *Instance) SetSoftware(software string) {
	i.Software = software
}
