package maps

// Keys returns keys of the map.
func Keys[M ~map[K]V, R []K, K comparable, V any](vals M) R {
	if len(vals) == 0 {
		return nil
	}

	out := make(R, 0, len(vals))
	for key := range vals {
		out = append(out, key)
	}
	return out
}
