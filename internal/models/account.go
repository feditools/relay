package models

import (
	"context"
	"github.com/uptrace/bun"
	"time"
)

// Account represents a federated social account.
type Account struct {
	ID          int64     `validate:"-" bun:"id,pk,autoincrement"`
	CreatedAt   time.Time `validate:"-" bun:",nullzero,notnull,default:current_timestamp"`
	UpdatedAt   time.Time `validate:"-" bun:",nullzero,notnull,default:current_timestamp"`
	ActorURI    string    `validate:"url" bun:",nullzero,notnull"`
	Username    string    `validate:"-" bun:",unique:unique_fedi_user,nullzero,notnull"`
	InstanceID  int64     `validate:"-" bun:",unique:unique_fedi_user,nullzero,notnull"`
	Instance    *Instance `validate:"-" bun:"rel:belongs-to,join:instance_id=id"`
	DisplayName string    `validate:"-" bun:",nullzero"`
	LastFinger  time.Time `validate:"-" bun:",notnull"`
	LogInCount  int64     `validate:"-" bun:",notnull"`
	LogInLast   time.Time `validate:"-" bun:",nullzero"`
	IsAdmin     bool      `validate:"-" bun:",notnull"`
	IsCouncil   bool      `validate:"-" bun:",notnull"`
}

var _ bun.BeforeAppendModelHook = (*Account)(nil)

// BeforeAppendModel runs before a bun append operation.
func (a *Account) BeforeAppendModel(_ context.Context, query bun.Query) error {
	switch query.(type) {
	case *bun.InsertQuery:
		now := time.Now()
		a.CreatedAt = now
		a.UpdatedAt = now

		err := validate.Struct(a)
		if err != nil {
			return err
		}
	case *bun.UpdateQuery:
		a.UpdatedAt = time.Now()

		err := validate.Struct(a)
		if err != nil {
			return err
		}
	}

	return nil
}
