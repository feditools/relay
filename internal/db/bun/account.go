package bun

import (
	"context"
	"database/sql"
	"errors"
	libdatabase "github.com/feditools/go-lib/database"
	"github.com/feditools/relay/internal/db"
	"github.com/feditools/relay/internal/models"
	"github.com/uptrace/bun"
	"time"
)

// CountAccounts returns the number of federated social account.
func (c *Client) CountAccounts(ctx context.Context) (int64, db.Error) {
	metric := c.metrics.NewDBQuery("CountAccounts")

	count, err := newAccountQ(c.bun, (*models.Account)(nil)).Count(ctx)

	if err != nil {
		go metric.Done(true)

		return 0, c.bun.errProc(err)
	}

	go metric.Done(false)

	return int64(count), nil
}

// CountAccountsForInstance returns the number of federated social account for an instance.
func (c *Client) CountAccountsForInstance(ctx context.Context, instanceID int64) (int64, db.Error) {
	metric := c.metrics.NewDBQuery("CountAccountsForInstance")

	count, err := newAccountQ(c.bun, (*models.Account)(nil)).Where("instance_id = ?", instanceID).Count(ctx)
	if err != nil {
		go metric.Done(true)

		return 0, c.bun.errProc(err)
	}

	go metric.Done(false)

	return int64(count), nil
}

// CreateAccount stores the federated social account.
func (c *Client) CreateAccount(ctx context.Context, account *models.Account) db.Error {
	metric := c.metrics.NewDBQuery("CreateAccount")

	if err := create(ctx, c.bun, account); err != nil {
		go metric.Done(true)

		return c.bun.errProc(err)
	}

	go metric.Done(false)

	return nil
}

// IncAccountLoginCount updates the login count of a stored federated instance.
func (c *Client) IncAccountLoginCount(ctx context.Context, account *models.Account) db.Error {
	metric := c.metrics.NewDBQuery("IncAccountLoginCount")

	err := c.bun.RunInTx(ctx, &sql.TxOptions{}, func(ctx context.Context, tx bun.Tx) error {
		err := tx.NewSelect().Model(account).Where("id = ?", account.ID).Scan(ctx)
		if err != nil {
			return err
		}

		account.LogInCount++
		account.LogInLast = time.Now()

		_, err = tx.NewUpdate().Model(account).Where("id = ?", account.ID).Exec(ctx)
		if err != nil {
			return err
		}

		return err
	})
	if err != nil {
		go metric.Done(true)

		return c.bun.errProc(err)
	}

	go metric.Done(false)

	return nil
}

// ReadAccount returns one federated social account.
func (c *Client) ReadAccount(ctx context.Context, id int64) (*models.Account, db.Error) {
	metric := c.metrics.NewDBQuery("ReadAccount")

	account := new(models.Account)
	err := newAccountQ(c.bun, account).Where("id = ?", id).Scan(ctx)
	if err != nil {
		dbErr := c.bun.ProcessError(err)

		if errors.Is(dbErr, db.ErrNoEntries) {
			// report no entries as a non error
			go metric.Done(false)
		} else {
			go metric.Done(true)
		}

		return nil, dbErr
	}

	go metric.Done(false)

	return account, nil
}

// ReadAccountByUsername returns one federated social account.
func (c *Client) ReadAccountByUsername(ctx context.Context, instanceID int64, username string) (*models.Account, db.Error) {
	metric := c.metrics.NewDBQuery("ReadAccountByUsername")

	account := new(models.Account)
	err := newAccountQ(c.bun, account).
		ColumnExpr("account.*").
		Join("RIGHT JOIN instances").
		JoinOn("account.instance_id = instances.id").
		Where("instances.id = ?", instanceID).
		Where("lower(account.username) = lower(?)", username).
		Scan(ctx)
	if err != nil {
		dbErr := c.bun.ProcessError(err)

		if errors.Is(dbErr, db.ErrNoEntries) {
			// report no entries as a non error
			go metric.Done(false)
		} else {
			go metric.Done(true)
		}

		return nil, dbErr
	}

	go metric.Done(false)

	return account, nil
}

// ReadAccountsPage returns a page of federated social accounts.
func (c *Client) ReadAccountsPage(ctx context.Context, index, count int) ([]*models.Account, db.Error) {
	metric := c.metrics.NewDBQuery("ReadAccountsPage")

	var accounts []*models.Account
	err := newAccountsQ(c.bun, &accounts).
		Limit(count).
		Offset(libdatabase.Offset(index, count)).
		Scan(ctx)
	if err != nil {
		go metric.Done(true)

		return nil, c.bun.ProcessError(err)
	}

	go metric.Done(false)

	return accounts, nil
}

// UpdateAccount updates the stored federated social account.
func (c *Client) UpdateAccount(ctx context.Context, account *models.Account) db.Error {
	metric := c.metrics.NewDBQuery("UpdateAccount")

	if err := update(ctx, c.bun, account); err != nil {
		go metric.Done(true)

		return c.bun.errProc(err)
	}

	go metric.Done(false)

	return nil
}

func newAccountQ(c bun.IDB, account *models.Account) *bun.SelectQuery {
	return c.
		NewSelect().
		Model(account)
}

func newAccountsQ(c bun.IDB, accounts *[]*models.Account) *bun.SelectQuery {
	return c.
		NewSelect().
		Model(accounts)
}
