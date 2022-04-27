package logic

import (
	"github.com/feditools/relay/internal/db"
	"sync"
	"time"
)

// Logic contains shared logic for the application
type Logic struct {
	db db.DB

	// peer list
	cPeerList         *[]string
	cPeerListExpires  time.Time
	cPeerListValidity time.Duration
	cPeerListLock     sync.RWMutex
}

// New created a new logic module
func New(d db.DB) (*Logic, error) {
	return &Logic{
		db: d,

		cPeerList:         &[]string{},
		cPeerListValidity: time.Second * 15,
	}, nil
}
