package bun

import (
	"errors"
	"fmt"
	"github.com/jackc/pgconn"
	"testing"
)

func TestProcessPostgresError(t *testing.T) {
	tables := []struct {
		x error
		n string
	}{
		{errors.New("test"), "test"},
		{&pgconn.PgError{Severity: "ERROR", Message: "random", Code: "12345"}, "ERROR: random (SQLSTATE 12345)"},
		{&pgconn.PgError{Severity: "ERROR", Message: "unique_violation", Code: "23505"}, "unique_violation"},
	}

	for i, table := range tables {
		i := i
		table := table

		name := fmt.Sprintf("[%d] Running processPostgresError for %s", i, table.x.Error())
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			err := processPostgresError(table.x)
			if err.Error() != table.n {
				t.Errorf("[%d] invalid error, got: '%s', want: '%s'", i, err.Error(), table.n)
			}
		})
	}
}

func TestProcessSQLiteError(t *testing.T) {
	tables := []struct {
		x error
		n string
	}{
		{errors.New("test"), "test"},
	}

	for i, table := range tables {
		i := i
		table := table

		name := fmt.Sprintf("[%d] Running processPostgresError for %s", i, table.x.Error())
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			err := processSQLiteError(table.x)
			if err.Error() != table.n {
				t.Errorf("[%d] invalid error, got: '%s', want: '%s'", i, err.Error(), table.n)
			}
		})
	}
}
