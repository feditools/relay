package logic

import (
	"fmt"
	"testing"
)

func TestTopDomains(t *testing.T) {
	tables := []struct {
		domain     string
		topDomains []string
	}{
		{"", []string{}},
		{"test.example.com", []string{"example.com", "com"}},
		{"super.cali.fragilistic.expi.ali.docious", []string{"cali.fragilistic.expi.ali.docious", "fragilistic.expi.ali.docious", "expi.ali.docious", "ali.docious", "docious"}},
	}

	for i, table := range tables {
		i := i
		table := table

		name := fmt.Sprintf("[%d] Running topDomains on %v", i, table.domain)
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			topDomains := topDomains(table.domain)
			if len(topDomains) != len(table.topDomains) {
				t.Errorf("[%d] invalid number of top domains, got: %d %v, want: %d %v", i, len(topDomains), topDomains, len(table.topDomains), table.topDomains)
				return
			}

			for j := 0; j < len(topDomains); j++ {
				if topDomains[j] != table.topDomains[j] {
					t.Errorf("[%d] invalid top domain at %d, got: '%s', want: '%s'", i, j, topDomains[j], table.topDomains[j])
				}
			}
		})
	}
}
