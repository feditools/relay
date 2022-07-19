package logic

import (
	"context"
	"github.com/feditools/relay/internal/config"
	"github.com/feditools/relay/internal/db"
	ihttp "github.com/feditools/relay/internal/http"
	"github.com/feditools/relay/internal/path"
	"github.com/feditools/relay/internal/runner"
	"github.com/feditools/relay/internal/transport"
	"github.com/go-fed/activity/pub"
	"github.com/go-fed/httpsig"
	lru "github.com/hashicorp/golang-lru"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"golang.org/x/sync/singleflight"
	"sync"
	"time"
)

// Logic contains shared logic for the application
type Logic struct {
	db        db.DB
	runner    runner.Runner
	transport *transport.Transport

	domain    string
	validAlgs []httpsig.Algorithm

	outgoingRequestGroup singleflight.Group

	// peer list
	cPeerList         *[]string
	cPeerListExpires  time.Time
	cPeerListValidity time.Duration
	cPeerListLock     sync.RWMutex

	// caches
	cacheActivity *lru.Cache
	cacheActor    *lru.Cache
	cacheDigest   *lru.Cache
}

// New created a new logic module
func New(ctx context.Context, c pub.Clock, d db.DB, h *ihttp.Client) (*Logic, error) {
	log := logger.WithFields(logrus.Fields{
		"func": "New",
	})

	// create module
	l := Logic{
		db: d,

		domain: viper.GetString(config.Keys.ServerExternalHostname),
		validAlgs: []httpsig.Algorithm{
			httpsig.RSA_SHA512,
			httpsig.RSA_SHA256,
			httpsig.ED25519,
		},

		cPeerList:         &[]string{},
		cPeerListValidity: time.Second * 15,
	}

	// get self
	instanceSelf, err := l.getOrCreateSelfInstance(ctx)
	if err != nil {
		log.Errorf("unable to get self: %s", err.Error())

		return nil, err
	}

	// generate transport
	l.transport, err = transport.New(c, h, path.GenPublicKey(l.domain), instanceSelf.PrivateKey)
	if err != nil {
		log.Errorf("creating transport: %s", err.Error())

		return nil, err
	}

	// make caches
	l.cacheActivity, err = lru.New(viper.GetInt(config.Keys.CachedActivityLimit))
	if err != nil {
		log.Errorf("creating activity cache: %s", err.Error())

		return nil, err
	}
	l.cacheActor, err = lru.New(viper.GetInt(config.Keys.CachedActorLimit))
	if err != nil {
		log.Errorf("creating actor cache: %s", err.Error())

		return nil, err
	}
	l.cacheDigest, err = lru.New(viper.GetInt(config.Keys.CachedDigestLimit))
	if err != nil {
		log.Errorf("creating digest cache: %s", err.Error())

		return nil, err
	}

	return &l, nil
}

func (l *Logic) SetRunner(r runner.Runner) {
	l.runner = r
}
