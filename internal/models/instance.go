package models

import (
	"context"
	"crypto/rsa"
	"github.com/uptrace/bun"
	"time"
)

// Instance represents a federated social instance
type Instance struct {
	ID         int64           `validate:"required" bun:",pk,nullzero,notnull,unique"`
	CreatedAt  time.Time       `validate:"-" bun:",nullzero,notnull,default:current_timestamp"`
	UpdatedAt  time.Time       `validate:"-" bun:",nullzero,notnull,default:current_timestamp"`
	Domain     string          `validate:"required,fqdn" bun:",nullzero,notnull,unique"`
	PrivateKey *rsa.PrivateKey `validate:"-"`
	PublicKey  *rsa.PublicKey  `validate:"-"`
	InboxIRI   string          `validate:"required,url" bun:",nullzero,notnull,unique"`
}

var _ bun.BeforeAppendModelHook = (*Instance)(nil)

// BeforeAppendModel runs before a bun append operation
func (f *Instance) BeforeAppendModel(_ context.Context, query bun.Query) error {
	switch query.(type) {
	case *bun.InsertQuery:
		now := time.Now()
		f.CreatedAt = now
		f.UpdatedAt = now

		err := validate.Struct(f)
		if err != nil {
			return err
		}
	case *bun.UpdateQuery:
		f.UpdatedAt = time.Now()

		err := validate.Struct(f)
		if err != nil {
			return err
		}
	}
	return nil
}
