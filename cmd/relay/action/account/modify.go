package account

import (
	"context"
	"github.com/feditools/go-lib"
	"github.com/feditools/go-lib/metrics/statsd"
	"github.com/feditools/relay/cmd/relay/action"
	"github.com/feditools/relay/internal/config"
	"github.com/feditools/relay/internal/db/bun"
	"github.com/spf13/viper"
)

// Modify runs database migrations.
var Modify action.Action = func(ctx context.Context) error {
	l := logger.WithField("func", "Modify")

	// create metrics collector
	metricsCollector, err := statsd.New(
		viper.GetString(config.Keys.MetricsStatsDAddress),
		viper.GetString(config.Keys.MetricsStatsDPrefix),
	)
	if err != nil {
		l.Errorf("metrics: %s", err.Error())

		return err
	}
	defer func() {
		l.Info("closing metrics collector")
		err := metricsCollector.Close()
		if err != nil {
			l.Errorf("closing metrics: %s", err.Error())
		}
	}()

	// create database client
	l.Info("creating database client")
	dbClient, err := bun.New(ctx, metricsCollector)
	if err != nil {
		l.Errorf("db: %s", err.Error())

		return err
	}
	defer func() {
		l.Info("closing database client")
		err := dbClient.Close(ctx)
		if err != nil {
			l.Errorf("closing db: %s", err.Error())
		}
	}()

	accountString := viper.GetString(config.Keys.Account)

	username, domain, err := lib.SplitAccount(accountString)
	if err != nil {
		l.Errorf("invalid account %s: %s", accountString, err.Error())

		return err
	}

	// find instance
	instance, err := dbClient.ReadInstanceByDomain(ctx, domain)
	if err != nil {
		l.Errorf("db read %s: %s", domain, err.Error())

		return err
	}
	if instance == nil {
		l.Infof("can't find instance %s", domain)

		return nil
	}
	l.Debugf("found instance %d: %+v", instance.ID, instance)

	// find account
	account, err := dbClient.ReadAccountByUsername(ctx, instance.ID, username)
	if err != nil {
		l.Errorf("db read %s: %s", username, err.Error())

		return err
	}
	if instance == nil {
		l.Infof("can't find user %s", username)

		return nil
	}
	l.Debugf("found account %d: %+v", account.ID, account)

	// add groups
	for _, addGroup := range viper.GetStringSlice(config.Keys.AccountAddGroup) {
		switch addGroup {
		case "admin":
			account.IsAdmin = true
		default:
			l.Warnf("unknown group %s, skipping", addGroup)
		}
	}

	// update database
	err = dbClient.UpdateAccount(ctx, account)
	if err != nil {
		l.Errorf("db update: %s", err.Error())

		return err
	}

	return nil
}
