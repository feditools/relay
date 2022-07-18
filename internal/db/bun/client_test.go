package bun

import (
	"context"
	"crypto/tls"
	"database/sql"
	"fmt"
	"github.com/feditools/go-lib/mock"
	"github.com/feditools/relay/internal/config"
	"github.com/feditools/relay/internal/db"
	"github.com/jackc/pgconn"
	"github.com/spf13/viper"
	"testing"
)

func TestDeriveBunDBPGOptions(t *testing.T) {
	dbDatabase := "database"
	dbPassword := "password"
	dbPort := 5432
	dbUser := "account"

	viper.Reset()

	viper.Set(config.Keys.DbType, "postgres")

	viper.Set(config.Keys.DbDatabase, dbDatabase)
	viper.Set(config.Keys.DbPassword, dbPassword)
	viper.Set(config.Keys.DbPort, dbPort)
	viper.Set(config.Keys.DbUser, dbUser)

	opts, err := deriveBunDBPGOptions()
	if err != nil {
		t.Errorf("unexpected error initializing pg options: %s", err.Error())
		return
	}
	if opts == nil {
		t.Errorf("opts is nil")
		return
	}

	if opts.Database != dbDatabase {
		t.Errorf("unexpected value for database, got: '%s', want: '%s'", opts.Database, dbDatabase)
	}
	if opts.Password != dbPassword {
		t.Errorf("unexpected value for password, got: '%s', want: '%s'", opts.Password, dbPassword)
	}
	if opts.Port != uint16(dbPort) {
		t.Errorf("unexpected value for port, got: '%d', want: '%d'", opts.Port, dbPort)
	}
	if opts.User != dbUser {
		t.Errorf("unexpected value for account, got: '%s', want: '%s'", opts.User, dbUser)
	}

	// tls
	if opts.TLSConfig != nil {
		t.Errorf("unexpected value for tls config, got: '%v', want: '%v'", opts.User, nil)
	}
}

func TestDeriveBunDBPGOptions_TLSDisable(t *testing.T) {
	dbAddress := "db.examle.com"
	dbDatabase := "database"
	dbPassword := "password"
	dbPort := 5432
	dbTLSMode := dbTLSModeDisable
	dbUser := "account"

	viper.Reset()

	viper.Set(config.Keys.DbType, "postgres")

	viper.Set(config.Keys.DbAddress, dbAddress)
	viper.Set(config.Keys.DbDatabase, dbDatabase)
	viper.Set(config.Keys.DbPassword, dbPassword)
	viper.Set(config.Keys.DbPort, dbPort)
	viper.Set(config.Keys.DbTLSMode, dbTLSMode)
	viper.Set(config.Keys.DbUser, dbUser)

	opts, err := deriveBunDBPGOptions()
	if err != nil {
		t.Errorf("unexpected error initializing pg options: %s", err.Error())
		return
	}
	if opts == nil {
		t.Errorf("opts is nil")
		return
	}

	if opts.Host != dbAddress {
		t.Errorf("unexpected value for address, got: '%s', want: '%s'", opts.Host, dbAddress)
	}
	if opts.Database != dbDatabase {
		t.Errorf("unexpected value for database, got: '%s', want: '%s'", opts.Database, dbDatabase)
	}
	if opts.Password != dbPassword {
		t.Errorf("unexpected value for password, got: '%s', want: '%s'", opts.Password, dbPassword)
	}
	if opts.Port != uint16(dbPort) {
		t.Errorf("unexpected value for port, got: '%d', want: '%d'", opts.Port, dbPort)
	}
	if opts.User != dbUser {
		t.Errorf("unexpected value for account, got: '%s', want: '%s'", opts.User, dbUser)
	}

	// tls
	if opts.TLSConfig != nil {
		t.Errorf("unexpected value for tls config, got: '%v', want: '%v'", opts.User, nil)
	}
}

func TestDeriveBunDBPGOptions_TLSEnable(t *testing.T) {
	dbAddress := "db.examle.com"
	dbDatabase := "database"
	dbPassword := "password"
	dbPort := 5432
	dbTLSMode := dbTLSModeEnable
	dbUser := "account"

	viper.Reset()

	viper.Set(config.Keys.DbType, "postgres")

	viper.Set(config.Keys.DbAddress, dbAddress)
	viper.Set(config.Keys.DbDatabase, dbDatabase)
	viper.Set(config.Keys.DbPassword, dbPassword)
	viper.Set(config.Keys.DbPort, dbPort)
	viper.Set(config.Keys.DbTLSMode, dbTLSMode)
	viper.Set(config.Keys.DbUser, dbUser)

	opts, err := deriveBunDBPGOptions()
	if err != nil {
		t.Errorf("unexpected error initializing pg options: %s", err.Error())
		return
	}
	if opts == nil {
		t.Errorf("opts is nil")
		return
	}

	if opts.Host != dbAddress {
		t.Errorf("unexpected value for address, got: '%s', want: '%s'", opts.Host, dbAddress)
	}
	if opts.Database != dbDatabase {
		t.Errorf("unexpected value for database, got: '%s', want: '%s'", opts.Database, dbDatabase)
	}
	if opts.Password != dbPassword {
		t.Errorf("unexpected value for password, got: '%s', want: '%s'", opts.Password, dbPassword)
	}
	if opts.Port != uint16(dbPort) {
		t.Errorf("unexpected value for port, got: '%d', want: '%d'", opts.Port, dbPort)
	}
	if opts.User != dbUser {
		t.Errorf("unexpected value for account, got: '%s', want: '%s'", opts.User, dbUser)
	}

	// tls
	if opts.TLSConfig == nil {
		t.Errorf("unexpected value for tls config, got: 'nil', want: '*tls.Config'")
		return
	}
	if opts.TLSConfig.InsecureSkipVerify != true {
		t.Errorf("unexpected value for tls inscure skip verfy, got: '%v', want: '%v'", opts.TLSConfig.InsecureSkipVerify, true)
	}
}

func TestDeriveBunDBPGOptions_TLSRequire(t *testing.T) {
	dbAddress := "db.examle.com"
	dbDatabase := "database"
	dbPassword := "password"
	dbPort := 5432
	dbTLSMode := dbTLSModeRequire
	dbUser := "account"

	viper.Reset()

	viper.Set(config.Keys.DbType, "postgres")

	viper.Set(config.Keys.DbAddress, dbAddress)
	viper.Set(config.Keys.DbDatabase, dbDatabase)
	viper.Set(config.Keys.DbPassword, dbPassword)
	viper.Set(config.Keys.DbPort, dbPort)
	viper.Set(config.Keys.DbTLSMode, dbTLSMode)
	viper.Set(config.Keys.DbUser, dbUser)

	viper.Set(config.Keys.DbTLSCACert, "../../../test/certificate.pem")

	opts, err := deriveBunDBPGOptions()
	if err != nil {
		t.Errorf("unexpected error initializing pg options: %s", err.Error())
		return
	}
	if opts == nil {
		t.Errorf("opts is nil")
		return
	}

	if opts.Host != dbAddress {
		t.Errorf("unexpected value for address, got: '%s', want: '%s'", opts.Host, dbAddress)
	}
	if opts.Database != dbDatabase {
		t.Errorf("unexpected value for database, got: '%s', want: '%s'", opts.Database, dbDatabase)
	}
	if opts.Password != dbPassword {
		t.Errorf("unexpected value for password, got: '%s', want: '%s'", opts.Password, dbPassword)
	}
	if opts.Port != uint16(dbPort) {
		t.Errorf("unexpected value for port, got: '%d', want: '%d'", opts.Port, dbPort)
	}
	if opts.User != dbUser {
		t.Errorf("unexpected value for account, got: '%s', want: '%s'", opts.User, dbUser)
	}

	// tls
	if opts.TLSConfig == nil {
		t.Errorf("unexpected value for tls config, got: 'nil', want: '*tls.Config'")
		return
	}
	if opts.TLSConfig.InsecureSkipVerify != false {
		t.Errorf("unexpected value for tls inscure skip verfy, got: '%v', want: '%v'", opts.TLSConfig.InsecureSkipVerify, false)
	}
	if opts.TLSConfig.ServerName != dbAddress {
		t.Errorf("unexpected value for tls inscure skip verfy, got: '%s', want: '%s'", opts.TLSConfig.ServerName, dbAddress)
	}
	if opts.TLSConfig.MinVersion != tls.VersionTLS12 {
		t.Errorf("unexpected value for tls inscure skip verfy, got: '%v', want: '%v'", opts.TLSConfig.MinVersion, tls.VersionTLS12)
	}
}

func TestDeriveBunDBPGOptions_NoDatabase(t *testing.T) {
	viper.Reset()

	viper.Set(config.Keys.DbType, "postgres")

	_, err := deriveBunDBPGOptions()
	errText := "no database set"
	if err.Error() != errText {
		t.Errorf("unexpected error initializing sqlite connection, got: '%s', want: '%s'", err.Error(), errText)
		return
	}
}

func TestDeriveBunDBPGOptions_NoType(t *testing.T) {
	viper.Reset()

	_, err := deriveBunDBPGOptions()
	errText := "expected bun type of POSTGRES but got "
	if err.Error() != errText {
		t.Errorf("unexpected error initializing sqlite connection, got: '%s', want: '%s'", err.Error(), errText)
		return
	}
}

func TestNew_Invalid(t *testing.T) {
	viper.Reset()

	viper.Set(config.Keys.DbType, "invalid")

	metricsCollector, _ := mock.NewMetricsCollector()

	_, err := New(context.Background(), metricsCollector)
	errText := "database type invalid not supported for bundb"
	if err.Error() != errText {
		t.Errorf("unexpected error initializing sqlite connection, got: '%s', want: '%s'", err.Error(), errText)
		return
	}
}

func TestNew_Sqlite(t *testing.T) {
	viper.Reset()

	viper.Set(config.Keys.DbType, "sqlite")
	viper.Set(config.Keys.DbAddress, ":memory:")

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

func TestPgConn_NoConfig(t *testing.T) {
	viper.Reset()

	_, err := pgConn(context.Background())
	errText := "could not create bundb postgres options: expected bun type of POSTGRES but got "
	if err.Error() != errText {
		t.Errorf("unexpected error initializing pg connection, got: '%s', want: '%s'", err.Error(), errText)
		return
	}
}

func TestSqliteConn(t *testing.T) {
	viper.Reset()

	viper.Set(config.Keys.DbAddress, ":memory:")

	bun, err := sqliteConn(context.Background())
	if err != nil {
		t.Errorf("unexpected error initializing sqlite connection: %s", err.Error())
		return
	}
	if bun == nil {
		t.Errorf("client is nil")
		return
	}
}

func TestSqliteConn_NoConfig(t *testing.T) {
	viper.Reset()

	_, err := sqliteConn(context.Background())
	errText := fmt.Sprintf("'%s' was not set when attempting to start sqlite", config.Keys.DbAddress)
	if err.Error() != errText {
		t.Errorf("unexpected error initializing sqlite connection, got: '%s', want: '%s'", err.Error(), errText)
		return
	}
}

func TestSqliteConn_BadPath(t *testing.T) {
	viper.Reset()

	viper.Set(config.Keys.DbAddress, "invalidir/db.sqlite")

	_, err := sqliteConn(context.Background())
	errText := "sqlite ping: Unable to open the database file (SQLITE_CANTOPEN)"
	if err.Error() != errText {
		t.Errorf("unexpected error initializing sqlite connection, got: '%s', want: '%s'", err.Error(), errText)
		return
	}
}

func TestProcessError(t *testing.T) {
	bun := &Bun{
		errProc: processPostgresError,
	}

	tables := []struct {
		x error
		n error
	}{
		{nil, nil},
		{sql.ErrNoRows, db.ErrNoEntries},
		{&pgconn.PgError{Severity: "ERROR", Message: "unique_violation", Code: "23505"}, db.NewErrAlreadyExists("unique_violation")},
	}

	for i, table := range tables {
		i := i
		table := table

		name := fmt.Sprintf("[%d] Running processPostgresError for %v", i, table.x)
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			err := bun.ProcessError(table.x)
			if table.x != nil {
				if err.Error() != table.n.Error() {
					t.Errorf("[%d] invalid error, got: '%s', want: '%s'", i, err.Error(), table.n.Error())
				}
			} else {
				if err != nil {
					t.Errorf("[%d] invalid error, got: '%s', want: 'nil'", i, err.Error())
				}
			}
		})
	}
}

func testNewSqliteClient() (db.DB, error) {
	viper.Reset()

	viper.Set(config.Keys.DbType, "sqlite")
	viper.Set(config.Keys.DbAddress, ":memory:")

	metricsCollector, _ := mock.NewMetricsCollector()

	client, err := New(context.Background(), metricsCollector)
	if err != nil {
		return nil, err
	}

	err = client.DoMigration(context.Background())
	if err != nil {
		return nil, err
	}

	if err != nil {
		return nil, err
	}

	return client, nil
}
