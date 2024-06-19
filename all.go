package maps

// All returns true if all key-value pairs match the given predicate.
func All[M ~map[K]V, K comparable, V any](vals M, predicate func(K, V) bool) bool {
	for key, val := range vals {
		if !predicate(key, val) {
			return false
		}
	}
	return true
}
