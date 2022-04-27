package bun

import (
	"context"
	"database/sql"
	"github.com/feditools/relay/internal/db"
	"github.com/feditools/relay/internal/models"
	"github.com/uptrace/bun"
	"time"
)

// CreateInstance stores the federated instance
func (c *Client) CreateInstance(ctx context.Context, instance *models.Instance) db.Error {
	start := time.Now()

	err := c.Create(ctx, instance)
	if err != nil {
		ended := time.Since(start)
		go c.metrics.DBQuery(ended, "CreateInstance", true)
		return c.bun.errProc(err)
	}

	ended := time.Since(start)
	go c.metrics.DBQuery(ended, "CreateInstance", false)
	return nil
}

// ReadInstanceByID returns one federated social instance
func (c *Client) ReadInstanceByID(ctx context.Context, id int64) (*models.Instance, db.Error) {
	start := time.Now()

	instance := &models.Instance{}

	err := c.newInstanceQ(instance).Where("id = ?", id).Scan(ctx)
	if err == sql.ErrNoRows {
		ended := time.Since(start)
		go c.metrics.DBQuery(ended, "ReadInstanceByID", false)
		return nil, nil
	}
	if err != nil {
		ended := time.Since(start)
		go c.metrics.DBQuery(ended, "ReadInstanceByID", true)
		return nil, c.bun.ProcessError(err)
	}

	ended := time.Since(start)
	go c.metrics.DBQuery(ended, "ReadInstanceByID", false)
	return instance, nil
}

// ReadInstanceByDomain returns one federated social instance
func (c *Client) ReadInstanceByDomain(ctx context.Context, domain string) (*models.Instance, db.Error) {
	start := time.Now()

	instance := &models.Instance{}

	err := c.newInstanceQ(instance).Where("lower(domain) = lower(?)", domain).Scan(ctx)
	if err == sql.ErrNoRows {
		ended := time.Since(start)
		go c.metrics.DBQuery(ended, "ReadInstanceByDomain", false)
		return nil, nil
	}
	if err != nil {
		ended := time.Since(start)
		go c.metrics.DBQuery(ended, "ReadInstanceByDomain", true)
		return nil, c.bun.ProcessError(err)
	}

	ended := time.Since(start)
	go c.metrics.DBQuery(ended, "ReadInstanceByDomain", false)
	return instance, nil
}

// UpdateInstance updates the stored federated instance
func (c *Client) UpdateInstance(ctx context.Context, instance *models.Instance) db.Error {
	start := time.Now()

	err := c.Update(ctx, instance)
	if err != nil {
		ended := time.Since(start)
		go c.metrics.DBQuery(ended, "UpdateInstance", true)
		return c.bun.errProc(err)
	}

	ended := time.Since(start)
	go c.metrics.DBQuery(ended, "UpdateInstance", false)
	return nil
}

func (c *Client) newInstanceQ(instance *models.Instance) *bun.SelectQuery {
	return c.bun.
		NewSelect().
		Model(instance)
}
