package models

import (
	"context"
	"crypto/rsa"
	"github.com/uptrace/bun"
	"time"
)

// Instance represents a federated social instance
type Instance struct {
	ID         int64           `validate:"-" bun:",pk,autoincrement,nullzero,notnull,unique"`
	CreatedAt  time.Time       `validate:"-" bun:",nullzero,notnull,default:current_timestamp"`
	UpdatedAt  time.Time       `validate:"-" bun:",nullzero,notnull,default:current_timestamp"`
	Domain     string          `validate:"required,fqdn" bun:",nullzero,notnull,unique"`
	PublicKey  *rsa.PublicKey  `validate:"-"`
	PrivateKey *rsa.PrivateKey `validate:"-"`
	ActorIRI   string          `validate:"required,url" bun:",nullzero,notnull,unique"`
	InboxIRI   string          `validate:"required,url" bun:",nullzero,notnull,unique"`
	Followed   bool            `validate:"-" bun:",notnull,default:false"`
	BlockID    int64           `validate:"-" bun:",nullzero"`
	Block      *Block          `validate:"-" bun:"rel:belongs-to"`
}

var _ bun.BeforeAppendModelHook = (*Instance)(nil)

// BeforeAppendModel runs before a bun append operation
func (i *Instance) BeforeAppendModel(_ context.Context, query bun.Query) error {
	switch query.(type) {
	case *bun.InsertQuery:
		now := time.Now()
		i.CreatedAt = now
		i.UpdatedAt = now

		err := validate.Struct(i)
		if err != nil {
			return err
		}
	case *bun.UpdateQuery:
		i.UpdatedAt = time.Now()

		err := validate.Struct(i)
		if err != nil {
			return err
		}
	}
	return nil
}
