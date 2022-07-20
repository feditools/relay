package webapp

import (
	"context"
	"encoding/gob"
	"github.com/feditools/go-lib/language"
	"github.com/feditools/go-lib/metrics"
	libtemplate "github.com/feditools/go-lib/template"
	"github.com/feditools/relay/internal/config"
	"github.com/feditools/relay/internal/db"
	"github.com/feditools/relay/internal/fedi"
	ihttp "github.com/feditools/relay/internal/http"
	itemplate "github.com/feditools/relay/internal/http/template"
	"github.com/feditools/relay/internal/kv"
	"github.com/feditools/relay/internal/logic/logic1"
	"github.com/feditools/relay/internal/path"
	"github.com/go-redis/redis/v8"
	"github.com/gorilla/sessions"
	"github.com/rbcervilla/redisstore/v8"
	"github.com/spf13/viper"
	"github.com/tdewolff/minify/v2"
	"github.com/tdewolff/minify/v2/html"
	"html/template"
	"strings"
	"sync"
	"time"
)

const SessionMaxAge = 30 * 24 * time.Hour // 30 days

// Module contains a webapp module for the web server. Implements web.Module.
type Module struct {
	db        db.DB
	fedi      *fedi.Module
	logic     *logic1.Logic
	language  *language.Module
	metrics   metrics.Collector
	minify    *minify.M
	store     sessions.Store
	srv       *ihttp.Server
	templates *template.Template

	domain        string
	logoSrcDark   string
	logoSrcLight  string
	headLinks     []libtemplate.HeadLink
	footerScripts []libtemplate.Script

	sigCache     map[string]string
	sigCacheLock sync.RWMutex
}

//revive:disable:argument-limit
// New returns a new webapp module.
func New(ctx context.Context, d db.DB, f *fedi.Module, lMod *language.Module, l *logic1.Logic, mc metrics.Collector, r redis.UniversalClient) (*Module, error) {
	log := logger.WithField("func", "New")

	// create new store
	store, err := redisstore.NewRedisStore(ctx, r)
	if err != nil {
		log.Errorf("create redis store: %s", err.Error())

		return nil, err
	}

	store.KeyPrefix(kv.KeySession())
	store.Options(sessions.Options{
		Path:   "/",
		Domain: viper.GetString(config.Keys.ServerExternalHostname),
		MaxAge: int(SessionMaxAge.Seconds()),
	})

	// register models for GOB
	gob.Register(SessionKey(0))

	// minify
	var m *minify.M
	if viper.GetBool(config.Keys.ServerMinifyHTML) {
		m = minify.New()
		m.AddFunc("text/html", html.Minify)
	}

	// get templates
	tmpl, err := itemplate.New()
	if err != nil {
		log.Errorf("create temates: %s", err.Error())

		return nil, err
	}

	// generate head links
	hl := []libtemplate.HeadLink{
		{
			HRef:        viper.GetString(config.Keys.WebappBootstrapCSSURI),
			Rel:         "stylesheet",
			CrossOrigin: COAnonymous,
			Integrity:   viper.GetString(config.Keys.WebappBootstrapCSSIntegrity),
		},
		{
			HRef:        viper.GetString(config.Keys.WebappFontAwesomeCSSURI),
			Rel:         "stylesheet",
			CrossOrigin: COAnonymous,
			Integrity:   viper.GetString(config.Keys.WebappFontAwesomeCSSIntegrity),
		},
	}
	paths := []string{
		path.FileDefaultCSS,
	}
	for _, p := range paths {
		signature, err := getSignature(strings.TrimPrefix(p, "/"))
		if err != nil {
			log.Errorf("getting signature for %s: %s", p, err.Error())
		}

		hl = append(hl, libtemplate.HeadLink{
			HRef:        p,
			Rel:         "stylesheet",
			CrossOrigin: COAnonymous,
			Integrity:   signature,
		})
	}

	// generate head links
	fs := []libtemplate.Script{
		{
			Src:         viper.GetString(config.Keys.WebappBootstrapJSURI),
			CrossOrigin: COAnonymous,
			Integrity:   viper.GetString(config.Keys.WebappBootstrapJSIntegrity),
		},
	}

	return &Module{
		db:        d,
		fedi:      f,
		language:  lMod,
		logic:     l,
		metrics:   mc,
		minify:    m,
		store:     store,
		templates: tmpl,

		domain:        viper.GetString(config.Keys.ServerExternalHostname),
		logoSrcDark:   viper.GetString(config.Keys.WebappLogoSrcDark),
		logoSrcLight:  viper.GetString(config.Keys.WebappLogoSrcLight),
		headLinks:     hl,
		footerScripts: fs,

		sigCache: map[string]string{},
	}, nil
} //revive:enable:argument-limit

// Name return the module name.
func (*Module) Name() string {
	return config.ServerRoleWebapp
}

// SetServer adds a reference to the server to the module.
func (m *Module) SetServer(s *ihttp.Server) {
	m.srv = s
}
