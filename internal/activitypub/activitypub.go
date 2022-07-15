package activitypub

import (
	"bytes"
	"context"
	"crypto/x509"
	"encoding/pem"
	"github.com/feditools/relay/internal/config"
	"github.com/feditools/relay/internal/db"
	"github.com/feditools/relay/internal/logic"
	"github.com/feditools/relay/internal/models"
	"github.com/feditools/relay/internal/path"
	"github.com/feditools/relay/internal/transport"
	"github.com/go-fed/activity/pub"
	"github.com/spf13/viper"
	"golang.org/x/sync/singleflight"
	"io"
	"sync"
)

// Module is an http module that handles activity pub activity
type Module struct {
	clock     pub.Clock
	db        db.DB
	logic     *logic.Logic
	transport *transport.Transport

	outgoingRequestGroup singleflight.Group

	appName      string
	appVersion   string
	domain       string
	publicKeyPem string

	inboxChan chan verifiedActivity
	inboxStop chan struct{}
	inboxWG   sync.WaitGroup
}

// New creates a new activity pub module
func New(ctx context.Context, d db.DB, c pub.Clock, l *logic.Logic) (*Module, error) {
	log := logger.WithField("func", "New")

	module := &Module{
		clock: c,
		db:    d,
		logic: l,

		appName:    viper.GetString(config.Keys.ApplicationName),
		appVersion: viper.GetString(config.Keys.SoftwareVersion),

		domain: viper.GetString(config.Keys.ServerExternalHostname),

		inboxChan: make(chan verifiedActivity, viper.GetInt(config.Keys.APInboxQueueSize)),
		inboxStop: make(chan struct{}),
	}

	var instanceSelf *models.Instance
	var err error
	instanceSelf, err = d.ReadInstanceByDomain(ctx, module.domain)
	if err != nil {
		log.Errorf("db read: %s", err.Error())
		return nil, err
	}

	if instanceSelf == nil {
		instanceSelf, err = module.createInstanceSelf(ctx)
		if err != nil {
			log.Errorf("create self: %s", err.Error())
			return nil, err
		}
	}

	// generate transport
	module.transport, err = transport.New(c, path.GenPublicKey(module.domain), instanceSelf.PrivateKey)
	if err != nil {
		log.Errorf("creating transport: %s", err.Error())
		return nil, err
	}

	// make public key pem
	publicKeyBytes, err := x509.MarshalPKIXPublicKey(instanceSelf.PublicKey)
	if err != nil {
		log.Errorf("marshaling public key: %s", err.Error())
		return nil, err
	}
	publicKeyBlock := &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: publicKeyBytes,
	}
	publicPem := new(bytes.Buffer)
	err = pem.Encode(publicPem, publicKeyBlock)
	if err != nil {
		log.Errorf("encoding pem: %s", err.Error())
		return nil, err
	}
	publicPemBytes, err := io.ReadAll(publicPem)
	if err != nil {
		log.Errorf("reading pem: %s", err.Error())
		return nil, err
	}
	module.publicKeyPem = string(publicPemBytes)

	for i := 0; i < viper.GetInt(config.Keys.APInboxWorkers); i++ {
		module.inboxWG.Add(1)
		go module.worker(ctx, i, &module.inboxWG, module.inboxChan, module.inboxStop)
	}

	return module, nil
}

// Close closes the activity pub module
func (m *Module) Close() {
	close(m.inboxStop)
}

// Name return the module name
func (m *Module) Name() string {
	return config.ServerRoleActivityPub
}
