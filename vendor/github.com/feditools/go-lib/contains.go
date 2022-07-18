package lib

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
