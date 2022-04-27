package logic

import "context"

// GetPeers returns true if a domain matches a block in the database
func (l *Logic) GetPeers(ctx context.Context) (*[]string, error) {
	log := logger.WithField("func", "GetPeers")

	instances, err := l.db.ReadInstancesWhereFollowing(ctx)
	if err != nil {
		log.Errorf("db read: %s", err.Error())
		return nil, err
	}
	log.Debugf("got %d instances", len(instances))

	// populate peer list
	peers := make([]string, len(instances))
	for i, instance := range instances {
		peers[i] = instance.Domain
	}

	return &peers, nil
}
