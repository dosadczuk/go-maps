package maps

// FilterByKey returns key-value pairs matching the given key predicate.
func FilterByKey[M ~map[K]V, R map[K]V, K comparable, V any](vals M, predicate func(K) bool) R {
	if len(vals) == 0 {
		return nil
	}

	out := make(R)
	for key, val := range vals {
		if predicate(key) {
			out[key] = val
		}
	}
	return out
}
