package maps

// MapValues returns a map of key-value pairs where values are created by applying
// transform function to each of the given pairs.
func MapValues[M ~map[K]V1, R map[K]V2, K comparable, V1, V2 any](vals M, transform func(K, V1) V2) R {
	if len(vals) == 0 {
		return nil
	}

	out := make(R, len(vals))
	for key, val := range vals {
		out[key] = transform(key, val)
	}
	return out
}
