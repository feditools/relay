//go:build postgres

package bun

import (
	"context"
	"github.com/feditools/relay/internal/config"
	"github.com/feditools/relay/internal/db"
	"github.com/feditools/relay/internal/mock"
	"github.com/feditools/relay/internal/models/testdata"
	"github.com/spf13/viper"
	"testing"
)

func TestNew_Postgres(t *testing.T) {
	dbAddress := "postgres"
	dbDatabase := "test"
	dbPassword := "test"
	dbPort := 5432
	dbTLSMode := dbTLSModeDisable
	dbUser := "test"

	viper.Reset()

	viper.Set(config.Keys.DbType, "postgres")

	viper.Set(config.Keys.DbAddress, dbAddress)
	viper.Set(config.Keys.DbDatabase, dbDatabase)
	viper.Set(config.Keys.DbPassword, dbPassword)
	viper.Set(config.Keys.DbPort, dbPort)
	viper.Set(config.Keys.DbTLSMode, dbTLSMode)
	viper.Set(config.Keys.DbUser, dbUser)

	metricsCollector, _ := mock.NewMetricsCollector()

	bun, err := New(context.Background(), metricsCollector)
	if err != nil {
		t.Errorf("unexpected error initializing bun connection: %s", err.Error())
		return
	}
	if bun == nil {
		t.Errorf("client is nil")
		return
	}
}

func TestPgConn(t *testing.T) {
	dbAddress := "postgres"
	dbDatabase := "test"
	dbPassword := "test"
	dbPort := 5432
	dbTLSMode := dbTLSModeDisable
	dbUser := "test"

	viper.Reset()

	viper.Set(config.Keys.DbType, "postgres")

	viper.Set(config.Keys.DbAddress, dbAddress)
	viper.Set(config.Keys.DbDatabase, dbDatabase)
	viper.Set(config.Keys.DbPassword, dbPassword)
	viper.Set(config.Keys.DbPort, dbPort)
	viper.Set(config.Keys.DbTLSMode, dbTLSMode)
	viper.Set(config.Keys.DbUser, dbUser)

	bun, err := pgConn(context.Background())
	if err != nil {
		t.Errorf("unexpected error initializing pg connection: %s", err.Error())
		return
	}
	if bun == nil {
		t.Errorf("client is nil")
		return
	}
}

func testNewPostresClient() (db.DB, error) {
	viper.Reset()

	viper.Set(config.Keys.DbType, "postgres")

	viper.Set(config.Keys.DbAddress, "postgres")
	viper.Set(config.Keys.DbDatabase, "test")
	viper.Set(config.Keys.DbPassword, "test")
	viper.Set(config.Keys.DbPort, 5432)
	viper.Set(config.Keys.DbUser, "test")
	viper.Set(config.Keys.DbEncryptionKey, testdata.TestEncryptionKey)

	metricsCollector, _ := mock.NewMetricsCollector()

	client, err := New(context.Background(), metricsCollector)
	if err != nil {
		return nil, err
	}

	err = client.DoMigration(context.Background())
	if err != nil {
		return nil, err
	}

	err = client.LoadTestData(context.Background())
	if err != nil {
		return nil, err
	}

	return client, nil
}
