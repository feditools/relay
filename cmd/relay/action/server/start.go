package server

import (
	"context"
	"github.com/feditools/relay/cmd/relay/action"
	"os"
	"os/signal"
	"syscall"
)

// Start starts the server
var Start action.Action = func(ctx context.Context) error {
	l := logger.WithField("func", "Start")

	l.Infof("starting")

	// ** start application **
	errChan := make(chan error)

	// Wait for SIGINT and SIGTERM (HIT CTRL-C)
	stopSigChan := make(chan os.Signal)
	signal.Notify(stopSigChan, syscall.SIGINT, syscall.SIGTERM)

	select {
	case sig := <-stopSigChan:
		l.Infof("got sig: %s", sig)
	case err := <-errChan:
		l.Fatal(err.Error())
	}

	l.Infof("done")
	return nil
}
