package bun

import (
	"context"
	"errors"
	"github.com/feditools/relay/internal/db"
	"github.com/feditools/relay/internal/models"
	"github.com/uptrace/bun"
)

// CreateBlock stores the domain block
func (c *Client) CreateBlock(ctx context.Context, block *models.Block) db.Error {
	metric := c.metrics.NewDBQuery("CreateBlock")

	if err := create(ctx, c.bun, block); err != nil {
		go metric.Done(true)

		return c.bun.errProc(err)
	}

	go metric.Done(false)

	return nil
}

// ReadBlock returns one domain block
func (c *Client) ReadBlock(ctx context.Context, id int64) (*models.Block, db.Error) {
	metric := c.metrics.NewDBQuery("ReadBlock")

	block := new(models.Block)
	err := newBlockQ(c.bun, block).Where("id = ?", id).Scan(ctx)
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

	return block, nil
}

// ReadBlockByDomain returns one domain block
func (c *Client) ReadBlockByDomain(ctx context.Context, domain string) (*models.Block, db.Error) {
	metric := c.metrics.NewDBQuery("ReadBlockByDomain")

	block := new(models.Block)
	err := newBlockQ(c.bun, block).Where("lower(domain) = lower(?)", domain).Scan(ctx)
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

	return block, nil
}

// UpdateBlock updates the stored domain block
func (c *Client) UpdateBlock(ctx context.Context, block *models.Block) db.Error {
	metric := c.metrics.NewDBQuery("UpdateBlock")

	if err := update(ctx, c.bun, block); err != nil {
		go metric.Done(true)

		return c.bun.errProc(err)
	}

	go metric.Done(false)

	return nil
}

func newBlockQ(c bun.IDB, block *models.Block) *bun.SelectQuery {
	return c.
		NewSelect().
		Model(block)
}
