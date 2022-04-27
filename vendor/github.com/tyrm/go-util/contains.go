package util

import "github.com/google/uuid"

// ContainsString will return true if a string is found in a given group of strings.
func ContainsString(stack []string, needle string) bool {
	for _, s := range stack {
		if s == needle {
			return true
		}
	}
	return false
}

// ContainsOneOfStrings will return true if any of a group of strings is found in a given group of strings.
func ContainsOneOfStrings(stack []string, needles []string) bool {
	for _, n := range needles {
		for _, s := range stack {
			if s == n {
				return true
			}
		}
	}
	return false
}

// ContainsUUID will return true if a uuid is found in a given group of uuids.
func ContainsUUID(stack []uuid.UUID, needle uuid.UUID) bool {
	for _, s := range stack {
		if s == needle {
			return true
		}
	}
	return false
}

// ContainsOneOfUUIDs will return true if any of a group of uuids is found in a given group of uuids.
func ContainsOneOfUUIDs(stack []uuid.UUID, needles []uuid.UUID) bool {
	for _, n := range needles {
		for _, s := range stack {
			if s == n {
				return true
			}
		}
	}
	return false
}
