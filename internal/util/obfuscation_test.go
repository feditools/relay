package util

import (
	"fmt"
	"testing"
)

func TestValidateDomainObfuscation(t *testing.T) {
	tables := []struct {
		domain            string
		obfuscationDomain string
		valid             bool
	}{
		{"example.com", "example.com", true},
		{"example.com", "e*****e.com", true},
		{"example.com", "e*******com", false},
		{"example.com", "example2.com", false},
		{"example.com", "peanuts.com", false},
	}

	for i, table := range tables {
		i := i
		table := table

		name := fmt.Sprintf("[%d] checking domain obfuscation", i)
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			if valid := ValidateDomainObfuscation(table.domain, table.obfuscationDomain); valid != table.valid {
				t.Errorf("wrong validation value, got: '%t', want: '%t'", valid, table.valid)
			}
		})
	}
}
