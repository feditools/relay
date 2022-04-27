package models

import "time"

// Block represents a block of a domain
type Block struct {
	ID              int64     `validate:"-" bun:",pk,autoincrement,nullzero,notnull,unique"`
	CreatedAt       time.Time `validate:"-" bun:",nullzero,notnull,default:current_timestamp"`
	UpdatedAt       time.Time `validate:"-" bun:",nullzero,notnull,default:current_timestamp"`
	Domain          string    `validate:"required,fqdn" bun:",nullzero,notnull,unique"`
	BlockSubdomains bool      `validate:"-" bun:",notnull,default:true"`
}
