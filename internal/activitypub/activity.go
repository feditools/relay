package activitypub

import (
	"fmt"
	"github.com/google/uuid"
)

func (m *Module) genActivityID() string {
	return fmt.Sprintf("https://%s/activities/%s", m.domain, uuid.New().String())
}

func (m *Module) genActorSelf() string {
	return fmt.Sprintf("https://%s/actor", m.domain)
}

func (m *Module) genActivityAccept(to, activityID string) map[string]interface{} {
	return map[string]interface{}{
		"@context": ContextActivityStreams,
		"type":     TypeAccept,
		"to":       []string{to},
		"actor":    m.genActorSelf(),
		"object": map[string]interface{}{
			"type":   TypeFollow,
			"id":     activityID,
			"object": m.genActorSelf(),
			"actor":  to,
		},
		"id": m.genActivityID(),
	}
}
