package bun

import (
	"github.com/feditools/relay/internal/db"
	"github.com/jackc/pgconn"
	"modernc.org/sqlite"
	sqlite3 "modernc.org/sqlite/lib"
)

// processPostgresError processes an error, replacing any postgres specific errors with our own error type
func processPostgresError(err error) db.Error {
	l := logger.WithField("func", "processPostgresError")

	// Attempt to cast as postgres
	pgErr, ok := err.(*pgconn.PgError)
	if !ok {
		return err
	}

	l.Debugf("postgres error %s: %s", pgErr.Code, pgErr.Error())

	// Handle supplied error code:
	// (https://www.postgresql.org/docs/10/errcodes-appendix.html)
	switch pgErr.Code {
	case "23505" /* unique_violation */ :
		return db.NewErrAlreadyExists(pgErr.Message)
	default:
		return err
	}
}

// processSQLiteError processes an error, replacing any sqlite specific errors with our own error type
func processSQLiteError(err error) db.Error {
	l := logger.WithField("func", "processSQLiteError")

	// Attempt to cast as sqlite
	sqliteErr, ok := err.(*sqlite.Error)
	if !ok {
		return err
	}

	l.Debugf("sqlite error %d: %s", sqliteErr.Code(), sqliteErr.Error())

	// Handle supplied error code:
	switch sqliteErr.Code() {
	case sqlite3.SQLITE_CONSTRAINT_UNIQUE, sqlite3.SQLITE_CONSTRAINT_PRIMARYKEY:
		return db.NewErrAlreadyExists(err.Error())
	default:
		return err
	}
}
