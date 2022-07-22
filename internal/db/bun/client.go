package bun

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"database/sql"
	"encoding/pem"
	"errors"
	"fmt"
	"github.com/feditools/go-lib/metrics"
	"github.com/feditools/relay/internal/config"
	"github.com/feditools/relay/internal/db"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/stdlib"
	"github.com/spf13/viper"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/dialect/sqlitedialect"
	"github.com/uptrace/bun/extra/bundebug"
	"modernc.org/sqlite"
	"os"
	"runtime"
	"strings"
)

const (
	dbTypePostgres = "postgres"
	dbTypeSqlite   = "sqlite"

	dbTLSModeDisable = "disable"
	dbTLSModeEnable  = "enable"
	dbTLSModeRequire = "require"
	dbTLSModeUnset   = ""
)

// Bun represents a bun db connection and it's error handler
type Bun struct {
	*bun.DB

	errProc func(error) db.Error
}

// Client is a DB interface compatible client for Bun
type Client struct {
	bun     *Bun
	metrics metrics.Collector
}

// New creates a new bun database client
func New(ctx context.Context, m metrics.Collector) (*Client, error) {
	var newBun *Bun
	var err error
	dbType := strings.ToLower(viper.GetString(config.Keys.DbType))

	switch dbType {
	case dbTypePostgres:
		newBun, err = pgConn(ctx)
		if err != nil {
			return nil, err
		}
	case dbTypeSqlite:
		newBun, err = sqliteConn(ctx)
		if err != nil {
			return nil, err
		}
	default:
		return nil, fmt.Errorf("database type %s not supported for bundb", dbType)
	}

	if strings.ToUpper(viper.GetString(config.Keys.LogLevel)) == "debug" || strings.ToUpper(viper.GetString(config.Keys.LogLevel)) == "trace" {
		newBun.AddQueryHook(bundebug.NewQueryHook(bundebug.WithVerbose(true)))
	}

	return &Client{
		bun:     newBun,
		metrics: m,
	}, nil
}

func sqliteConn(ctx context.Context) (*Bun, error) {
	l := logger.WithField("func", "sqliteConn")

	// validate bun address has actually been set
	dbAddress := viper.GetString(config.Keys.DbAddress)
	if dbAddress == "" {
		return nil, fmt.Errorf("'%s' was not set when attempting to start sqlite", config.Keys.DbAddress)
	}

	// Drop anything fancy from DB address
	dbAddress = strings.Split(dbAddress, "?")[0]
	dbAddress = strings.TrimPrefix(dbAddress, "file:")

	// Append our own SQLite preferences
	dbAddress = "file:" + dbAddress + "?cache=shared"

	// Open new DB instance
	sqldb, err := sql.Open("sqlite", dbAddress)
	if err != nil {
		if errWithCode, ok := err.(*sqlite.Error); ok {
			err = errors.New(sqlite.ErrorCodeString[errWithCode.Code()])
		}
		return nil, fmt.Errorf("could not open sqlite bun: %s", err)
	}

	setConnectionValues(sqldb)

	if dbAddress == "file::memory:?cache=shared" {
		l.Warn("sqlite in-memory database should only be used for debugging")
		// don't close connections on disconnect -- otherwise
		// the SQLite database will be deleted when there
		// are no active connections
		sqldb.SetConnMaxLifetime(0)
	}

	conn := getErrConn(bun.NewDB(sqldb, sqlitedialect.New()))

	// ping to check the bun is there and listening
	if err := conn.PingContext(ctx); err != nil {
		if errWithCode, ok := err.(*sqlite.Error); ok {
			err = errors.New(sqlite.ErrorCodeString[errWithCode.Code()])
		}
		return nil, fmt.Errorf("sqlite ping: %s", err)
	}

	l.Info("connected to SQLITE database")
	return conn, nil
}

func pgConn(ctx context.Context) (*Bun, error) {
	l := logger.WithField("func", "pgConn")

	opts, err := deriveBunDBPGOptions()
	if err != nil {
		return nil, fmt.Errorf("could not create bundb postgres options: %s", err)
	}

	sqldb := stdlib.OpenDB(*opts)

	setConnectionValues(sqldb)

	conn := getErrConn(bun.NewDB(sqldb, pgdialect.New()))

	// ping to check the bun is there and listening
	if err := conn.PingContext(ctx); err != nil {
		return nil, fmt.Errorf("postgres ping: %s", err)
	}

	l.Info("connected to POSTGRES database")
	return conn, nil
}

func deriveBunDBPGOptions() (*pgx.ConnConfig, error) {
	keys := config.Keys

	if strings.ToUpper(viper.GetString(keys.DbType)) != db.TypePostgres {
		return nil, fmt.Errorf("expected bun type of %s but got %s", db.TypePostgres, viper.GetString(keys.DbType))
	}

	// these are all optional, the bun adapter figures out defaults
	port := viper.GetInt(keys.DbPort)
	address := viper.GetString(keys.DbAddress)
	username := viper.GetString(keys.DbUser)
	password := viper.GetString(keys.DbPassword)

	// validate database
	database := viper.GetString(keys.DbDatabase)
	if database == "" {
		return nil, errors.New("no database set")
	}

	var tlsConfig *tls.Config
	tlsMode := viper.GetString(keys.DbTLSMode)
	switch tlsMode {
	case dbTLSModeDisable, dbTLSModeUnset:
		break // nothing to do
	case dbTLSModeEnable:
		/* #nosec G402 */
		tlsConfig = &tls.Config{
			InsecureSkipVerify: true,
		}
	case dbTLSModeRequire:
		tlsConfig = &tls.Config{
			InsecureSkipVerify: false,
			ServerName:         viper.GetString(keys.DbAddress),
			MinVersion:         tls.VersionTLS12,
		}
	}

	caCertPath := viper.GetString(keys.DbTLSCACert)
	if tlsConfig != nil && caCertPath != "" {
		// load the system cert pool first -- we'll append the given CA cert to this
		certPool, err := x509.SystemCertPool()
		if err != nil {
			return nil, fmt.Errorf("error fetching system CA cert pool: %s", err)
		}

		// open the file itself and make sure there's something in it
		/* #nosec G304 */
		caCertBytes, err := os.ReadFile(caCertPath)
		if err != nil {
			return nil, fmt.Errorf("error opening CA certificate at %s: %s", caCertPath, err)
		}
		if len(caCertBytes) == 0 {
			return nil, fmt.Errorf("ca cert at %s was empty", caCertPath)
		}

		// make sure we have a PEM block
		caPem, _ := pem.Decode(caCertBytes)
		if caPem == nil {
			return nil, fmt.Errorf("could not parse cert at %s into PEM", caCertPath)
		}

		// parse the PEM block into the certificate
		caCert, err := x509.ParseCertificate(caPem.Bytes)
		if err != nil {
			return nil, fmt.Errorf("could not parse cert at %s into x509 certificate: %s", caCertPath, err)
		}

		// we're happy, add it to the existing pool and then use this pool in our tls config
		certPool.AddCert(caCert)
		tlsConfig.RootCAs = certPool
	}

	cfg, _ := pgx.ParseConfig("")
	if address != "" {
		cfg.Host = address
	}
	if port > 0 {
		cfg.Port = uint16(port)
	}

	if username != "" {
		cfg.User = username
	}
	if password != "" {
		cfg.Password = password
	}
	if tlsConfig != nil {
		cfg.TLSConfig = tlsConfig
	}
	cfg.Database = database
	cfg.PreferSimpleProtocol = true
	cfg.RuntimeParams["application_name"] = viper.GetString(keys.ApplicationName)

	return cfg, nil
}

// https://bun.uptrace.dev/postgres/running-bun-in-production.html#database-sql
func setConnectionValues(sqldb *sql.DB) {
	maxOpenConns := 4 * runtime.GOMAXPROCS(0)
	sqldb.SetMaxOpenConns(maxOpenConns)
	sqldb.SetMaxIdleConns(maxOpenConns)
}

func getErrConn(dbConn *bun.DB) *Bun {
	var errProc func(error) db.Error
	switch dbConn.Dialect().Name() {
	case dialect.PG:
		errProc = processPostgresError
	case dialect.SQLite:
		errProc = processSQLiteError
	default:
		panic("unknown dialect name: " + dbConn.Dialect().Name().String())
	}
	return &Bun{
		errProc: errProc,
		DB:      dbConn,
	}
}

// ProcessError replaces any known values with our own db.Error types
func (conn *Bun) ProcessError(err error) db.Error {
	switch {
	case err == nil:
		return nil
	case err == sql.ErrNoRows:
		return db.ErrNoEntries
	default:
		return conn.errProc(err)
	}
}
