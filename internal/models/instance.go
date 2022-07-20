package models

import (
	"bytes"
	"context"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"github.com/sirupsen/logrus"
	"github.com/uptrace/bun"
	"io"
	"time"
)

// Instance represents a federated social instance
type Instance struct {
	ID             int64           `validate:"-" bun:",pk,autoincrement,nullzero,notnull,unique"`
	CreatedAt      time.Time       `validate:"-" bun:",nullzero,notnull,default:current_timestamp"`
	UpdatedAt      time.Time       `validate:"-" bun:",nullzero,notnull,default:current_timestamp"`
	Domain         string          `validate:"required,fqdn" bun:",nullzero,notnull,unique"`
	ServerHostname string          `validate:"required,fqdn" bun:",nullzero,notnull,unique"`
	Software       string          `validate:"-" bun:",nullzero"`
	PublicKey      *rsa.PublicKey  `validate:"-"`
	PrivateKey     *rsa.PrivateKey `validate:"-"`
	ActorIRI       string          `validate:"required,url" bun:",nullzero,notnull,unique"`
	InboxIRI       string          `validate:"required,url" bun:",nullzero,notnull,unique"`
	Followed       bool            `validate:"-" bun:",notnull"`
	BlockID        int64           `validate:"-" bun:",nullzero"`
	Block          *Block          `validate:"-" bun:"rel:belongs-to"`
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

func (i *Instance) PublicKeyPEM() (string, error) {
	l := logger.WithFields(logrus.Fields{
		"func":  "PublicKeyPEM",
		"model": "Instance",
	})

	publicKeyBytes, err := x509.MarshalPKIXPublicKey(i.PublicKey)
	if err != nil {
		l.Errorf("marshaling public key: %s", err.Error())

		return "", err
	}
	publicKeyBlock := &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: publicKeyBytes,
	}
	publicPem := new(bytes.Buffer)
	err = pem.Encode(publicPem, publicKeyBlock)
	if err != nil {
		l.Errorf("encoding pem: %s", err.Error())

		return "", err
	}
	publicPemBytes, err := io.ReadAll(publicPem)
	if err != nil {
		l.Errorf("reading pem: %s", err.Error())

		return "", err
	}

	return string(publicPemBytes), nil
}
