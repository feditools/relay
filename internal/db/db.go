package db

import (
	"context"
	"github.com/feditools/relay/internal/models"
)

// DB represents a database client
type DB interface {
	// Close closes the db connections
	Close(ctx context.Context) Error
	// Create stores the object
	Create(ctx context.Context, i any) Error
	// DoMigration runs database migrations
	DoMigration(ctx context.Context) Error
	// LoadTestData adds test data to the database
	LoadTestData(ctx context.Context) Error
	// ReadByID returns a model by its ID
	ReadByID(ctx context.Context, id int64, i any) Error
	// Update updates stored data
	Update(ctx context.Context, i any) Error

	// Instance

	// CreateInstance stores the federated instance and caches it
	CreateInstance(ctx context.Context, instance *models.Instance) (err Error)
	// ReadInstanceByID returns one federated social instance.
	ReadInstanceByID(ctx context.Context, id int64) (instance *models.Instance, err Error)
	// ReadInstanceByDomain returns one federated social instance.
	ReadInstanceByDomain(ctx context.Context, domain string) (instance *models.Instance, err Error)
	// UpdateInstance updates the stored federated instance and caches it
	UpdateInstance(ctx context.Context, instance *models.Instance) (err Error)
}
