package maps

// FilterByValue returns key-value pairs matching the given value predicate.
func FilterByValue[M ~map[K]V, R map[K]V, K comparable, V any](vals M, predicate func(V) bool) R {
	if len(vals) == 0 {
		return nil
	}

	out := make(R)
	for key, val := range vals {
		if predicate(val) {
			out[key] = val
		}
	}
	return out
}
