package db

import (
	"fmt"
	"testing"
)

func TestNewErrAlreadyExists(t *testing.T) {
	tables := []struct {
		x string
		n bool
	}{
		{"test1", false},
		{"test2", false},
		{"test3", false},
		{"test4", false},
		{"test5", false},
	}

	for i, table := range tables {
		i := i
		table := table

		name := fmt.Sprintf("[%d] Running NewErrAlreadyExists with '%s'", i, table.x)
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			err := NewErrAlreadyExists(table.x)
			if err.Error() != table.x {
				t.Errorf("[%d] got wrong error text for '%s', got: '%s', want: '%s'", i, table.x, err.Error(), table.x)
			}

		})
	}
}
