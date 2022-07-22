package db

import (
	"context"
	"github.com/feditools/relay/internal/models"
)

// DB represents a database client
type DB interface {
	// Close closes the db connections
	Close(ctx context.Context) Error
	// DoMigration runs database migrations
	DoMigration(ctx context.Context) Error

	// Accounts

	// CountAccounts returns the number of federated social account
	CountAccounts(ctx context.Context) (count int64, err Error)
	// CountAccountsForInstance returns the number of federated social account for an instance
	CountAccountsForInstance(ctx context.Context, instanceID int64) (count int64, err Error)
	// CreateAccount stores the federated social account
	CreateAccount(ctx context.Context, account *models.Account) (err Error)
	// IncAccountLoginCount updates the login count of a stored federated instance
	IncAccountLoginCount(ctx context.Context, account *models.Account) (err Error)
	// ReadAccount returns one federated social account
	ReadAccount(ctx context.Context, id int64) (account *models.Account, err Error)
	// ReadAccountByUsername returns one federated social account
	ReadAccountByUsername(ctx context.Context, instanceID int64, username string) (account *models.Account, err Error)
	// ReadAccountsPage returns a page of federated social accounts
	ReadAccountsPage(ctx context.Context, index, count int) (instances []*models.Account, err Error)
	// UpdateAccount updates the stored federated instance
	UpdateAccount(ctx context.Context, account *models.Account) (err Error)

	// Block

	// CreateBlock stores the domain block
	CreateBlock(ctx context.Context, block *models.Block) (err Error)
	// DeleteBlock deletes a domain block
	DeleteBlock(ctx context.Context, block *models.Block) (err Error)
	// ReadBlock returns one domain block
	ReadBlock(ctx context.Context, id int64) (block *models.Block, err Error)
	// ReadBlockByDomain returns one domain block by domain name
	ReadBlockByDomain(ctx context.Context, domain string) (block *models.Block, err Error)
	// ReadBlocks returns all domain block
	ReadBlocks(ctx context.Context) (block []*models.Block, err Error)
	// ReadBlocksPage returns a page of domain blocks
	ReadBlocksPage(ctx context.Context, index, count int) (instances []*models.Block, err Error)
	// UpdateBlock updates the stored domain block
	UpdateBlock(ctx context.Context, block *models.Block) (err Error)

	// Instance

	// CreateInstance stores the federated instance
	CreateInstance(ctx context.Context, instance *models.Instance) (err Error)
	// ReadInstance returns one federated social instance
	ReadInstance(ctx context.Context, id int64) (instance *models.Instance, err Error)
	// ReadInstanceByActorIRI returns one federated social instance
	ReadInstanceByActorIRI(ctx context.Context, actorIRI string) (instance *models.Instance, err Error)
	// ReadInstanceByDomain returns one federated social instance
	ReadInstanceByDomain(ctx context.Context, domain string) (instance *models.Instance, err Error)
	// ReadInstancesWhereFollowing returns all federated social instances which are following this relay
	ReadInstancesWhereFollowing(ctx context.Context) (instances []*models.Instance, err Error)
	// UpdateInstance updates the stored federated instance
	UpdateInstance(ctx context.Context, instance *models.Instance) (err Error)
}
