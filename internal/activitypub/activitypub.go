package activitypub

import (
	"bytes"
	"context"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"github.com/feditools/relay/internal/config"
	"github.com/feditools/relay/internal/db"
	"github.com/feditools/relay/internal/models"
	"github.com/feditools/relay/internal/path"
	"github.com/spf13/viper"
	"io"
)

// Module is an http module that handles activity pub activity
type Module struct {
	db db.DB

	appName      string
	domain       string
	privateKey   *rsa.PrivateKey
	publicKey    *rsa.PublicKey
	publicKeyPem string
}

// New creates a new activity pub module
func New(ctx context.Context, d db.DB) (*Module, error) {
	l := logger.WithField("func", "New")

	module := &Module{
		db: d,

		appName: viper.GetString(config.Keys.ApplicationName),
		domain:  viper.GetString(config.Keys.ServerExternalHostname),
	}

	var instanceSelf *models.Instance
	var err error
	instanceSelf, err = d.ReadInstanceByDomain(ctx, module.domain)
	if err != nil {
		l.Errorf("db read: %s", err.Error())
		return nil, err
	}

	if instanceSelf == nil {
		instanceSelf, err = module.createInstanceSelf(ctx)
		if err != nil {
			l.Errorf("create self: %s", err.Error())
			return nil, err
		}
	}

	privateKey, err := instanceSelf.GetPrivateKey()
	if err != nil {
		l.Errorf("decrypting private key: %s", err.Error())
		return nil, err
	}

	// add keys
	module.privateKey = privateKey
	module.publicKey = instanceSelf.PublicKey

	// make public key pem
	publicKeyBytes, err := x509.MarshalPKIXPublicKey(instanceSelf.PublicKey)
	if err != nil {
		l.Errorf("marshaling public key: %s", err.Error())
		return nil, err
	}
	publicKeyBlock := &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: publicKeyBytes,
	}
	publicPem := new(bytes.Buffer)
	err = pem.Encode(publicPem, publicKeyBlock)
	if err != nil {
		l.Errorf("encoding pem: %s", err.Error())
		return nil, err
	}
	publicPemBytes, err := io.ReadAll(publicPem)
	if err != nil {
		l.Errorf("reading pem: %s", err.Error())
		return nil, err
	}
	module.publicKeyPem = string(publicPemBytes)

	return module, nil
}

func (m *Module) createInstanceSelf(ctx context.Context) (*models.Instance, error) {
	l := logger.WithField("func", "createInstanceSelf")

	// generate key
	privatekey, err := rsa.GenerateKey(rand.Reader, ActorKeySize)
	if err != nil {
		l.Errorf("genrating private key: %s", err.Error())
		return nil, err
	}

	publickey := &privatekey.PublicKey

	// create new instance
	newInstance := &models.Instance{
		Domain:   m.domain,
		InboxIRI: path.GenInbox(m.domain),

		PublicKey: publickey,
	}

	// set private key
	err = newInstance.SetPrivateKey(privatekey)
	if err != nil {
		l.Errorf("setting private key: %s", err.Error())
		return nil, err
	}

	// add to database
	err = m.db.CreateInstance(ctx, newInstance)
	if err != nil {
		l.Errorf("db create: %s", err.Error())
		return nil, err
	}

	return newInstance, nil
}

// Name return the module name
func (m *Module) Name() string {
	return config.ServerRoleActivityPub
}
