package helpers

// Clears the entire slice to save some memory and will be reused.
func ClearSlice[T any](s []T) []T {
	clear(s)
	return s[:0]
}
