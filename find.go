package maps

// Find returns the first value matching the given predicate.
func Find[M ~map[K]V, K comparable, V any](vals M, predicate func(K, V) bool) (out V, found bool) {
	for key, val := range vals {
		if predicate(key, val) {
			return val, true
		}
	}
	return
}
