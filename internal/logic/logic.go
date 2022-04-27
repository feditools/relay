package logic

import "github.com/feditools/relay/internal/db"

// Logic contains shared logic for the application
type Logic struct {
	db db.DB
}

// New created a new logic module
func New(d db.DB) (*Logic, error) {
	return &Logic{
		db: d,
	}, nil
}
