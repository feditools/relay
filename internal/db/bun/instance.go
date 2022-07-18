package bun

import (
	"context"
	"errors"
	"github.com/feditools/relay/internal/db"
	"github.com/feditools/relay/internal/models"
	"github.com/uptrace/bun"
)

// CreateInstance stores the federated instance
func (c *Client) CreateInstance(ctx context.Context, instance *models.Instance) db.Error {
	metric := c.metrics.NewDBQuery("CreateInstance")

	if err := create(ctx, c.bun, instance); err != nil {
		go metric.Done(true)

		return c.bun.errProc(err)
	}

	go metric.Done(false)

	return nil
}

// ReadInstance returns one federated social instance
func (c *Client) ReadInstance(ctx context.Context, id int64) (*models.Instance, db.Error) {
	metric := c.metrics.NewDBQuery("ReadInstance")

	instance := new(models.Instance)
	err := newInstanceQ(c.bun, instance).Where("id = ?", id).Scan(ctx)
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

	return instance, nil
}

// ReadInstanceByActorIRI returns one federated social instance
func (c *Client) ReadInstanceByActorIRI(ctx context.Context, actorIRI string) (*models.Instance, db.Error) {
	metric := c.metrics.NewDBQuery("ReadInstanceByActorIRI")

	instance := new(models.Instance)
	err := newInstanceQ(c.bun, instance).Where("actor_iri = ?", actorIRI).Scan(ctx)
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

	return instance, nil
}

// ReadInstanceByDomain returns one federated social instance
func (c *Client) ReadInstanceByDomain(ctx context.Context, domain string) (*models.Instance, db.Error) {
	metric := c.metrics.NewDBQuery("ReadInstanceByDomain")

	instance := new(models.Instance)
	err := newInstanceQ(c.bun, instance).Where("lower(domain) = lower(?)", domain).Scan(ctx)
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

	return instance, nil
}

// ReadInstancesWhereFollowing returns all federated social instances which are following this relay
func (c *Client) ReadInstancesWhereFollowing(ctx context.Context) ([]*models.Instance, db.Error) {
	metric := c.metrics.NewDBQuery("ReadInstancesWhereFollowing")

	var instances []*models.Instance
	err := newInstancesQ(c.bun, &instances).Where("followed = true").Scan(ctx)
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

	return instances, nil
}

// UpdateInstance updates the stored federated instance
func (c *Client) UpdateInstance(ctx context.Context, instance *models.Instance) db.Error {
	metric := c.metrics.NewDBQuery("UpdateInstance")

	if err := update(ctx, c.bun, instance); err != nil {
		go metric.Done(true)

		return c.bun.errProc(err)
	}

	go metric.Done(false)

	return nil
}

func newInstanceQ(c bun.IDB, instance *models.Instance) *bun.SelectQuery {
	return c.
		NewSelect().
		Model(instance)
}

func newInstancesQ(c bun.IDB, instances *[]*models.Instance) *bun.SelectQuery {
	return c.
		NewSelect().
		Model(instances)
}
