package bun

import (
	"context"
	"github.com/feditools/relay/internal/db"
	"github.com/feditools/relay/internal/db/bun/migrations"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/migrate"
)

// Close closes the bun db connection
func (c *Client) Close(ctx context.Context) db.Error {
	l := logger.WithField("func", "Close")
	l.Info("closing db connection")
	return c.bun.Close()
}

// DoMigration runs schema migrations on the database
func (c *Client) DoMigration(ctx context.Context) db.Error {
	l := logger.WithField("func", "DoMigration")

	migrator := migrate.NewMigrator(c.bun.DB, migrations.Migrations)

	if err := migrator.Init(ctx); err != nil {
		return err
	}

	group, err := migrator.Migrate(ctx)
	if err != nil {
		if err.Error() == "migrate: there are no any migrations" {
			return nil
		}
		return err
	}

	if group.ID == 0 {
		l.Info("there are no new migrations to run")
		return nil
	}

	l.Infof("migrated database to %s", group)
	return nil
}

func create(ctx context.Context, c bun.IDB, i interface{}) error {
	_, err := c.NewInsert().Model(i).Exec(ctx)
	if err != nil {
		logger.WithField("func", "create").Error(err.Error())
	}

	return err
}

func delete(ctx context.Context, c bun.IDB, i interface{}) error {
	_, err := c.NewDelete().Model(i).Exec(ctx)
	if err != nil {
		logger.WithField("func", "delete").Error(err.Error())
	}

	return err
}

func update(ctx context.Context, c bun.IDB, i interface{}) error {
	q := c.NewUpdate().Model(i).WherePK()

	_, err := q.Exec(ctx)
	if err != nil {
		logger.WithField("func", "update").Error(err.Error())
	}

	return err
}
