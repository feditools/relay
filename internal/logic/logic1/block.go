package logic1

import (
	"context"
	"errors"
	"github.com/feditools/relay/internal/db"
	"strings"
)

// IsDomainBlocked returns true if a domain matches a block in the database
func (l *Logic) IsDomainBlocked(ctx context.Context, d string) (bool, error) {
	log := logger.WithField("func", "IsDomainBlocked")

	// check domain for block
	_, err := l.db.ReadBlockByDomain(ctx, d)
	if err != nil && !errors.Is(err, db.ErrNoEntries) {
		log.Errorf("db read %s: %s", d, err.Error())

		return false, err
	}

	// this domain is blocked
	if err == nil {
		return true, nil
	}

	// check top domains
	for _, domain := range topDomains(d) {
		block, err := l.db.ReadBlockByDomain(ctx, domain)
		if err != nil && !errors.Is(err, db.ErrNoEntries) {
			log.Errorf("db read %s: %s", domain, err.Error())

			return false, err
		}
		if err == nil {
			if block.BlockSubdomains {
				return true, nil
			}
		}
	}

	return false, nil
}

func topDomains(d string) []string {
	parts := strings.Split(d, ".")
	end := len(parts)

	tds := make([]string, len(parts)-1)
	for i := 0; i < len(tds); i++ {
		tds[i] = strings.Join(parts[i+1:end], ".")
	}

	return tds
}
