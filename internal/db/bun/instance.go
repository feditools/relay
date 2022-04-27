package bun

import (
	"context"
	"database/sql"
	"github.com/feditools/relay/internal/db"
	"github.com/feditools/relay/internal/models"
	"github.com/uptrace/bun"
)

// CreateInstance stores the federated instance
func (c *Client) CreateInstance(ctx context.Context, instance *models.Instance) db.Error {
	err := c.Create(ctx, instance)
	if err != nil {
		return c.bun.errProc(err)
	}
	return nil
}

// ReadInstanceByID returns one federated social instance.
func (c *Client) ReadInstanceByID(ctx context.Context, id int64) (*models.Instance, db.Error) {
	instance := &models.Instance{}

	err := c.newFediInstanceQ(instance).Where("id = ?", id).Scan(ctx)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, c.bun.ProcessError(err)
	}

	return instance, nil
}

// ReadInstanceByDomain returns one federated social instance.
func (c *Client) ReadInstanceByDomain(ctx context.Context, domain string) (*models.Instance, db.Error) {
	instance := &models.Instance{}

	err := c.newFediInstanceQ(instance).Where("lower(domain) = lower(?)", domain).Scan(ctx)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, c.bun.ProcessError(err)
	}
	return instance, nil
}

// UpdateInstance updates the stored federated instance
func (c *Client) UpdateInstance(ctx context.Context, instance *models.Instance) db.Error {
	err := c.Update(ctx, instance)
	if err != nil {
		return c.bun.errProc(err)
	}
	return nil
}

func (c *Client) newFediInstanceQ(instance *models.Instance) *bun.SelectQuery {
	return c.bun.
		NewSelect().
		Model(instance)
}
