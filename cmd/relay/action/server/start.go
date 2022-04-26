package server

import (
	"context"
	"fmt"
	"github.com/feditools/relay/cmd/relay/action"
	"github.com/feditools/relay/internal/http"
	"os"
	"os/signal"
	"syscall"
)

// Start starts the server
var Start action.Action = func(ctx context.Context) error {
	l := logger.WithField("func", "Start")

	l.Info("starting")

	l.Debug("creating http server")
	server, err := http.NewServer(ctx)
	if err != nil {
		l.Errorf("http server: %s", err.Error())
		return err
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
	case err := <-errChan:
		l.Fatal(err.Error())
	}

	l.Infof("done")
	return nil
}
