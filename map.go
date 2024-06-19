package maps

// Map returns a slice of values created by applying transform function
// to each of the given key-value pairs.
func Map[M ~map[K]V, R []E, K comparable, V any, E any](vals M, transform func(K, V) E) R {
	if len(vals) == 0 {
		return nil
	}

	out := make(R, 0, len(vals))
	for key, val := range vals {
		out = append(out, transform(key, val))
	}
	return out
}
