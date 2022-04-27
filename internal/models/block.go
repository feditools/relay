package models

import (
	"context"
	"github.com/uptrace/bun"
	"time"
)

// Block represents a block of a domain
type Block struct {
	ID              int64     `validate:"-" bun:",pk,autoincrement,nullzero,notnull,unique"`
	CreatedAt       time.Time `validate:"-" bun:",nullzero,notnull,default:current_timestamp"`
	UpdatedAt       time.Time `validate:"-" bun:",nullzero,notnull,default:current_timestamp"`
	Domain          string    `validate:"required,fqdn" bun:",nullzero,notnull,unique"`
	BlockSubdomains bool      `validate:"-" bun:",notnull,default:true"`
}

var _ bun.BeforeAppendModelHook = (*Block)(nil)

// BeforeAppendModel runs before a bun append operation
func (b *Block) BeforeAppendModel(_ context.Context, query bun.Query) error {
	switch query.(type) {
	case *bun.InsertQuery:
		now := time.Now()
		b.CreatedAt = now
		b.UpdatedAt = now

		err := validate.Struct(b)
		if err != nil {
			return err
		}
	case *bun.UpdateQuery:
		b.UpdatedAt = time.Now()

		err := validate.Struct(b)
		if err != nil {
			return err
		}
	}
	return nil
}
