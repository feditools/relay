package server

import (
	"context"
	"fmt"
	"github.com/feditools/go-lib"
	"github.com/feditools/go-lib/language"
	"github.com/feditools/go-lib/metrics/statsd"
	"github.com/feditools/relay/cmd/relay/action"
	"github.com/feditools/relay/internal/clock"
	"github.com/feditools/relay/internal/config"
	"github.com/feditools/relay/internal/db/bun"
	"github.com/feditools/relay/internal/fedi"
	"github.com/feditools/relay/internal/http"
	"github.com/feditools/relay/internal/http/activitypub"
	"github.com/feditools/relay/internal/http/static"
	"github.com/feditools/relay/internal/http/webapp"
	"github.com/feditools/relay/internal/kv/redis"
	"github.com/feditools/relay/internal/logic/logic1"
	"github.com/feditools/relay/internal/runner/faktory"
	"github.com/feditools/relay/internal/token"
	"github.com/spf13/viper"
	"os"
	"os/signal"
	"syscall"
)

// Start starts the server
var Start action.Action = func(topCtx context.Context) error {
	l := logger.WithField("func", "Start")
	l.Info("starting")

	ctx, cancel := context.WithCancel(topCtx)

	// create metrics collector
	metricsCollector, err := statsd.New(
		viper.GetString(config.Keys.MetricsStatsDAddress),
		viper.GetString(config.Keys.MetricsStatsDPrefix),
	)
	if err != nil {
		l.Errorf("metrics: %s", err.Error())
		cancel()

		return err
	}
	defer func() {
		l.Info("closing metrics collector")
		err := metricsCollector.Close()
		if err != nil {
			l.Errorf("closing metrics: %s", err.Error())
		}
	}()

	// create clock module
	l.Debug("creating clock")
	clockMod := clock.NewClock()

	// create database client
	l.Debug("creating database client")
	dbClient, err := bun.New(ctx, metricsCollector)
	if err != nil {
		l.Errorf("db: %s", err.Error())
		cancel()

		return err
	}
	defer func() {
		err := dbClient.Close(ctx)
		if err != nil {
			l.Errorf("closing db: %s", err.Error())
		}
	}()

	// create http client
	httpClient, err := http.NewClient(ctx)
	if err != nil {
		l.Errorf("http client: %s", err.Error())
		cancel()

		return err
	}

	// create kv client
	kvClient, err := redis.New(ctx)
	if err != nil {
		l.Errorf("redis: %s", err.Error())
		cancel()

		return err
	}
	defer func() {
		err := kvClient.Close(ctx)
		if err != nil {
			l.Errorf("closing redis: %s", err.Error())
		}
	}()

	// create language module
	languageMod, err := language.New()
	if err != nil {
		l.Errorf("language: %s", err.Error())
		cancel()

		return err
	}

	// create tokenizer
	tokz, err := token.New()
	if err != nil {
		l.Errorf("create tokenizer: %s", err.Error())
		cancel()

		return err
	}

	// create logic module
	l.Debug("creating logic module")
	logicMod, err := logic1.New(ctx, clockMod, dbClient, httpClient)
	if err != nil {
		l.Errorf("logic: %s", err.Error())
		cancel()

		return err
	}

	// create fedi module
	fediMod, err := fedi.New(dbClient, logicMod.Transport(), kvClient, tokz)
	if err != nil {
		l.Errorf("fedi: %s", err.Error())
		cancel()

		return err
	}
	logicMod.SetFedi(fediMod)

	// create runner
	runnerMod, err := faktory.New(logicMod)
	if err != nil {
		l.Errorf("runner: %s", err.Error())
		cancel()

		return err
	}
	logicMod.SetRunner(runnerMod)
	runnerMod.Start(ctx)

	// create http server
	l.Debug("creating http server")
	server, err := http.NewServer(ctx, metricsCollector)
	if err != nil {
		l.Errorf("http server: %s", err.Error())
		cancel()

		return err
	}

	// create web modules
	var webModules []http.Module
	if lib.ContainsString(viper.GetStringSlice(config.Keys.ServerRoles), config.ServerRoleActivityPub) {
		l.Infof("adding %s module", config.ServerRoleActivityPub)
		apMod, err := activitypub.New(ctx, logicMod, runnerMod)
		if err != nil {
			l.Errorf("%s: %s", config.ServerRoleActivityPub, err.Error())
			cancel()

			return err
		}
		webModules = append(webModules, apMod)
	}
	if lib.ContainsString(viper.GetStringSlice(config.Keys.ServerRoles), config.ServerRoleStatic) {
		l.Infof("adding %s module", config.ServerRoleStatic)
		apMod, err := static.New()
		if err != nil {
			l.Errorf("%s: %s", config.ServerRoleStatic, err.Error())
			cancel()

			return err
		}
		webModules = append(webModules, apMod)
	}
	if lib.ContainsString(viper.GetStringSlice(config.Keys.ServerRoles), config.ServerRoleWebapp) {
		l.Infof("adding %s module", config.ServerRoleWebapp)
		apMod, err := webapp.New(
			ctx,
			dbClient,
			fediMod,
			languageMod,
			logicMod,
			metricsCollector,
			kvClient.RedisClient(),
			runnerMod,
			tokz,
		)
		if err != nil {
			l.Errorf("%s: %s", config.ServerRoleWebapp, err.Error())
			cancel()

			return err
		}
		webModules = append(webModules, apMod)
	}

	// add modules to servers
	for _, mod := range webModules {
		mod.SetServer(server)
		err := mod.Route(server)
		if err != nil {
			l.Errorf("loading %s module: %s", mod.Name(), err.Error())
			cancel()

			return err
		}
	}

	// ** start application **
	errChan := make(chan error)

	// Wait for SIGINT and SIGTERM (HIT CTRL-C)
	stopSigChan := make(chan os.Signal)
	signal.Notify(stopSigChan, syscall.SIGINT, syscall.SIGTERM)

	// start webserver
	go func(s *http.Server, errChan chan error) {
		l.Debug("starting http server")
		err := s.Start()
		if err != nil {
			errChan <- fmt.Errorf("http server: %s", err.Error())
		}
	}(server, errChan)

	// wait for event
	select {
	case sig := <-stopSigChan:
		l.Infof("got sig: %s", sig)
		cancel()
	case err := <-errChan:
		l.Fatal(err.Error())
		cancel()
	}

	<-ctx.Done()
	l.Infof("done")
	return nil
}
