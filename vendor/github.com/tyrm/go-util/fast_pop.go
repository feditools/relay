package util

// FastPopString will remove the first found string from group of strings without preserving the order
func FastPopString(slice []string, elem string) []string {
	// look for elem in slice
	index := -1
	for i, n := range slice {
		if n == elem {
			index = i
		}
	}

	// elem not found
	if index == -1 {
		return slice
	}

	// element is last element
	if index != -1 && len(slice) == 1 {
		return []string{}
	}

	// Remove the element at index i from a.
	slice[index] = slice[len(slice)-1] // Copy last element to index i.
	slice[len(slice)-1] = ""           // Erase last element (write zero value).
	slice = slice[:len(slice)-1]       // Truncate slice.

	return slice
}
