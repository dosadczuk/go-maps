package maps

// Filter returns key-value pairs matching the given predicate.
func Filter[M ~map[K]V, R map[K]V, K comparable, V any](vals M, predicate func(K, V) bool) R {
	if len(vals) == 0 {
		return nil
	}

	out := make(R)
	for key, val := range vals {
		if predicate(key, val) {
			out[key] = val
		}
	}
	return out
}
