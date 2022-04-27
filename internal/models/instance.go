package models

import (
	"context"
	"crypto/rsa"
	"crypto/x509"
	"github.com/uptrace/bun"
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

// GetPrivateKey returns unencrypted private key
func (i *Instance) GetPrivateKey() (*rsa.PrivateKey, error) {
	// decrypt bytes
	b, err := decrypt(i.PrivateKey)

	// convert to private key
	pk, err := x509.ParsePKCS1PrivateKey(b)
	if err != nil {
		return nil, err
	}

	// return private key
	return pk, err
}

// SetPrivateKey sets encrypted private key
func (i *Instance) SetPrivateKey(pk *rsa.PrivateKey) error {
	// convert to bytes
	b := x509.MarshalPKCS1PrivateKey(pk)

	// encrypt bytes
	data, err := encrypt(b)
	if err != nil {
		return err
	}

	// store bytes
	i.PrivateKey = data
	return nil
}
