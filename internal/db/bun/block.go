package bun

import (
	"context"
	"database/sql"
	"github.com/feditools/relay/internal/db"
	"github.com/feditools/relay/internal/models"
	"github.com/uptrace/bun"
	"time"
)

// CreateBlock stores the domain block
func (c *Client) CreateBlock(ctx context.Context, block *models.Block) db.Error {
	start := time.Now()

	err := c.Create(ctx, block)
	if err != nil {
		ended := time.Since(start)
		go c.metrics.DBQuery(ended, "CreateBlock", true)
		return c.bun.errProc(err)
	}

	ended := time.Since(start)
	go c.metrics.DBQuery(ended, "CreateBlock", false)
	return nil
}

// ReadBlockByID returns one domain block
func (c *Client) ReadBlockByID(ctx context.Context, id int64) (*models.Block, db.Error) {
	start := time.Now()

	block := &models.Block{}

	err := c.newBlockQ(block).Where("id = ?", id).Scan(ctx)
	if err == sql.ErrNoRows {
		ended := time.Since(start)
		go c.metrics.DBQuery(ended, "ReadBlockByID", false)
		return nil, nil
	}
	if err != nil {
		ended := time.Since(start)
		go c.metrics.DBQuery(ended, "ReadBlockByID", true)
		return nil, c.bun.ProcessError(err)
	}

	ended := time.Since(start)
	go c.metrics.DBQuery(ended, "ReadBlockByID", false)
	return block, nil
}

// ReadBlockByDomain returns one domain block
func (c *Client) ReadBlockByDomain(ctx context.Context, domain string) (*models.Block, db.Error) {
	start := time.Now()

	block := &models.Block{}

	err := c.newBlockQ(block).Where("lower(domain) = lower(?)", domain).Scan(ctx)
	if err == sql.ErrNoRows {
		ended := time.Since(start)
		go c.metrics.DBQuery(ended, "ReadBlockByDomain", false)
		return nil, nil
	}
	if err != nil {
		ended := time.Since(start)
		go c.metrics.DBQuery(ended, "ReadBlockByDomain", true)
		return nil, c.bun.ProcessError(err)
	}

	ended := time.Since(start)
	go c.metrics.DBQuery(ended, "ReadBlockByDomain", false)
	return block, nil
}

// UpdateBlock updates the stored domain block
func (c *Client) UpdateBlock(ctx context.Context, block *models.Block) db.Error {
	start := time.Now()

	err := c.Update(ctx, block)
	if err != nil {
		ended := time.Since(start)
		go c.metrics.DBQuery(ended, "UpdateBlock", true)
		return c.bun.errProc(err)
	}

	ended := time.Since(start)
	go c.metrics.DBQuery(ended, "UpdateBlock", false)
	return nil
}

func (c *Client) newBlockQ(block *models.Block) *bun.SelectQuery {
	return c.bun.
		NewSelect().
		Model(block)
}
