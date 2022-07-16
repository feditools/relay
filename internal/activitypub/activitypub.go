package activitypub

import (
	"context"
	"github.com/feditools/relay/internal/config"
	"github.com/feditools/relay/internal/logic"
	"github.com/feditools/relay/internal/runner"
	"github.com/spf13/viper"
)

// Module is an http module that handles activity pub activity
type Module struct {
	logic  *logic.Logic
	runner runner.Runner

	appName      string
	appVersion   string
	publicKeyPem string
}

// New creates a new activity pub module
func New(ctx context.Context, l *logic.Logic, r runner.Runner) (*Module, error) {
	log := logger.WithField("func", "New")

	module := &Module{
		logic:  l,
		runner: r,

		appName:    viper.GetString(config.Keys.ApplicationName),
		appVersion: viper.GetString(config.Keys.SoftwareVersion),
	}

	// get self
	instanceSelf, err := module.logic.GetInstanceSelf(ctx)
	if err != nil {
		log.Errorf("getting self: %s", err.Error())
		return nil, err
	}

	// get public key pem
	publicKeyPem, err := instanceSelf.PublicKeyPEM()
	if err != nil {
		log.Errorf("getting instance pem: %s", err.Error())

		return nil, err
	}
	module.publicKeyPem = publicKeyPem

	return module, nil
}

// Name return the module name
func (m *Module) Name() string {
	return config.ServerRoleActivityPub
}
