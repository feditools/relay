package logic

import (
	"context"
	"time"
)

// GetPeers returns a list of peers
func (l *Logic) GetPeers(ctx context.Context) (*[]string, error) {
	log := logger.WithField("func", "GetPeers")

	// check cache
	peers, valid := l.getCachedPeerList()
	if valid {
		return peers, nil
	}

	// read from db
	instances, err := l.db.ReadInstancesWhereFollowing(ctx)
	if err != nil {
		log.Errorf("db read: %s", err.Error())
		return nil, err
	}
	log.Debugf("got %d instances", len(instances))

	// populate peer list
	newPeers := make([]string, len(instances))
	for i, instance := range instances {
		newPeers[i] = instance.Domain
	}

	// update cache
	l.setCachedPeerList(&newPeers)

	return &newPeers, nil
}

// peer list "cache" functions
func (l *Logic) getCachedPeerList() (*[]string, bool) {
	l.cPeerListLock.RLock()
	if l.cPeerListExpires.Before(time.Now()) {
		l.cPeerListLock.RUnlock()
		return nil, false
	}
	count := l.cPeerList
	l.cPeerListLock.RUnlock()
	return count, true
}

func (l *Logic) setCachedPeerList(peers *[]string) {
	l.cPeerListLock.Lock()
	l.cPeerList = peers
	l.cPeerListExpires = time.Now().Add(l.cPeerListValidity)
	l.cPeerListLock.Unlock()
}
