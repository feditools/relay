package models

import (
	"crypto/rsa"
	"time"
)

// Instance represents a federated social instance
type Instance struct {
	ID         int64          `validate:"-" bun:",pk,autoincrement,nullzero,notnull,unique"`
	CreatedAt  time.Time      `validate:"-" bun:",nullzero,notnull,default:current_timestamp"`
	UpdatedAt  time.Time      `validate:"-" bun:",nullzero,notnull,default:current_timestamp"`
	Domain     string         `validate:"required,fqdn" bun:",nullzero,notnull,unique"`
	PublicKey  *rsa.PublicKey `validate:"-"`
	PrivateKey []byte         `validate:"-"`
	InboxIRI   string         `validate:"required,url" bun:",nullzero,notnull,unique"`
	BlockID    int64          `validate:"-" bun:",nullzero"`
	Block      *Block         `validate:"-" bun:"rel:belongs-to"`
}
