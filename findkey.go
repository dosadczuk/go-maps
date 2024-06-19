package maps

// FindKey returns the first key matching the given predicate.
func FindKey[M ~map[K]V, K comparable, V any](vals M, predicate func(K, V) bool) (out K, found bool) {
	for key, val := range vals {
		if predicate(key, val) {
			return key, true
		}
	}
	return
}
