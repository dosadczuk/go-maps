package maps

// Any returns true if any of the key-value pairs match the given predicate.
// If predicate is `nil`, returns true if map has at least one element.
func Any[M ~map[K]V, K comparable, V any](vals M, predicate func(K, V) bool) bool {
	if predicate == nil {
		return len(vals) > 0
	}

	for key, val := range vals {
		if predicate(key, val) {
			return true
		}
	}
	return false
}
