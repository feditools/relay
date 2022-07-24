package logic1

import (
	"context"
	"errors"
	"github.com/feditools/relay/internal/db"
	"github.com/feditools/relay/internal/models"
)

func (l *Logic) AddBlock(ctx context.Context, block *models.Block) error {
	log := logger.WithField("func", "AddBlock")

	// add block to database
	err := l.db.CreateBlock(ctx, block)
	if err != nil {
		log.Errorf("db create: %s", err.Error())
		return err
	}

	// enqueue block update
	err = l.runner.EnqueueProcessBlock(ctx, block.ID)
	if err != nil {
		log.Errorf("enqueueing job: %s", err.Error())

		return err
	}

	return nil
}

// IsDomainBlocked returns true if a domain matches a block in the database
func (l *Logic) IsDomainBlocked(ctx context.Context, d string) (bool, error) {
	log := logger.WithField("func", "IsDomainBlocked")
	log.Tracef("checking for blocked domain %s", d)

	// check domain for block
	_, err := l.db.ReadBlockByDomain(ctx, d)
	if err != nil && !errors.Is(err, db.ErrNoEntries) {
		log.Errorf("db read %s: %s", d, err.Error())

		return false, err
	}

	// this domain is blocked
	if err == nil {
		log.Tracef("found blocked domain %s", d)

		return true, nil
	}
	log.Tracef("block not found for domain %s, checking subdomains", d)

	// check top domains
	for _, domain := range topDomains(d) {
		log.Tracef("checking for blocked top domain %s", domain)
		block, err := l.db.ReadBlockByDomain(ctx, domain)
		if err != nil && !errors.Is(err, db.ErrNoEntries) {
			log.Errorf("db read %s: %s", domain, err.Error())

			return false, err
		}
		if err == nil {
			log.Tracef("found block for top domain %s, include subdaomins: %t", d, block.BlockSubdomains)
			if block.BlockSubdomains {
				return true, nil
			}
		}
	}

	return false, nil
}

func (l *Logic) ProcessBlock(ctx context.Context, jid string, blockID int64) error {
	return nil
}
