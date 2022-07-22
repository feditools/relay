package logic1

import (
	"fmt"
	"strings"
)

func genActorSelf(domain string) string {
	return fmt.Sprintf("https://%s/actor", domain)
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
