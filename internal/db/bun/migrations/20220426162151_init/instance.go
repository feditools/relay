package models

import (
	"crypto/rsa"
	"time"
)

// Instance represents a federated social instance
type Instance struct {
	ID             int64           `validate:"-" bun:",pk,autoincrement,nullzero,notnull,unique"`
	CreatedAt      time.Time       `validate:"-" bun:",nullzero,notnull,default:current_timestamp"`
	UpdatedAt      time.Time       `validate:"-" bun:",nullzero,notnull,default:current_timestamp"`
	AccountDomain  string          `validate:"required,fqdn" bun:",nullzero,notnull,unique"`
	ServerHostname string          `validate:"required,fqdn" bun:",nullzero,notnull,unique"`
	Software       string          `validate:"-" bun:",nullzero"`
	PublicKey      *rsa.PublicKey  `validate:"-"`
	PrivateKey     *rsa.PrivateKey `validate:"-"`
	ActorIRI       string          `validate:"required,url" bun:",nullzero,notnull,unique"`
	InboxIRI       string          `validate:"required,url" bun:",nullzero,notnull,unique"`
	Followed       bool            `validate:"-" bun:",notnull,default:false"`
	BlockID        int64           `validate:"-" bun:",nullzero"`
	Block          *Block          `validate:"-" bun:"rel:belongs-to"`
}
