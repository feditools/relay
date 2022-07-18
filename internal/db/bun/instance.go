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
		return nil, c.bun.ProcessError(err)
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

// ReadInstanceByActorIRI returns one federated social instance
func (c *Client) ReadInstanceByActorIRI(ctx context.Context, actorIRI string) (*models.Instance, db.Error) {
	start := time.Now()

	instance := new(models.Instance)

	err := c.newInstanceQ(instance).Where("actor_iri = ?", actorIRI).Scan(ctx)
	if err == sql.ErrNoRows {
		ended := time.Since(start)
		go c.metrics.DBQuery(ended, "ReadInstanceByActorIRI", false)
		return nil, c.bun.ProcessError(err)
	}
	if err != nil {
		ended := time.Since(start)
		go c.metrics.DBQuery(ended, "ReadInstanceByActorIRI", true)
		return nil, c.bun.ProcessError(err)
	}

	ended := time.Since(start)
	go c.metrics.DBQuery(ended, "ReadInstanceByActorIRI", false)
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
		return nil, c.bun.ProcessError(err)
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

// ReadInstancesWhereFollowing returns all federated social instances which are following this relay
func (c *Client) ReadInstancesWhereFollowing(ctx context.Context) ([]*models.Instance, db.Error) {
	start := time.Now()

	var instances []*models.Instance

	err := c.newInstancesQ(&instances).Where("followed = true").Scan(ctx)
	if err != nil {
		ended := time.Since(start)
		go c.metrics.DBQuery(ended, "ReadInstancesWhereFollowing", true)
		return nil, c.bun.ProcessError(err)
	}

	ended := time.Since(start)
	go c.metrics.DBQuery(ended, "ReadInstancesWhereFollowing", false)
	return instances, nil
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

func (c *Client) newInstancesQ(instances *[]*models.Instance) *bun.SelectQuery {
	return c.bun.
		NewSelect().
		Model(instances)
}
