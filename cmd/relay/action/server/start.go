package server

import (
	"context"
	"fmt"
	"github.com/feditools/relay/cmd/relay/action"
	"github.com/feditools/relay/internal/activitypub"
	"github.com/feditools/relay/internal/clock"
	"github.com/feditools/relay/internal/config"
	"github.com/feditools/relay/internal/db/bun"
	"github.com/feditools/relay/internal/http"
	"github.com/feditools/relay/internal/logic"
	"github.com/feditools/relay/internal/metrics/statsd"
	"github.com/feditools/relay/internal/runner/faktory"
	"github.com/spf13/viper"
	"github.com/tyrm/go-util"
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
	metricsCollector, err := statsd.New()
	if err != nil {
		l.Errorf("metrics: %s", err.Error())
		cancel()

		return err
	}
	defer func() {
		err := metricsCollector.Close()
		if err != nil {
			l.Errorf("closing metrics: %s", err.Error())
		}
	}()

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

	// create clock module
	l.Debug("creating clock")
	clockMod := clock.NewClock()

	// create logic module
	l.Debug("creating logic module")
	logicMod, err := logic.New(ctx, clockMod, dbClient)
	if err != nil {
		l.Errorf("logic: %s", err.Error())
		cancel()

		return err
	}

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
	if util.ContainsString(viper.GetStringSlice(config.Keys.ServerRoles), config.ServerRoleActivityPub) {
		l.Infof("adding %s module", config.ServerRoleActivityPub)
		apMod, err := activitypub.New(ctx, logicMod, runnerMod)
		if err != nil {
			l.Errorf("%s: %s", config.ServerRoleActivityPub, err.Error())
			cancel()

			return err
		}
		webModules = append(webModules, apMod)
	}

	// add modules to servers
	for _, mod := range webModules {
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
