package database

// Offset is a tiny func that returns the offset for an index and count.
func Offset(i, c int) int { return i * c }
