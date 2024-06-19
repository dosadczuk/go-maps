package maps

// MapKeys returns a map of key-value pairs where keys are created by applying
// transform function to each of the given pairs.
func MapKeys[M ~map[K1]V, R map[K2]V, K1, K2 comparable, V any](vals M, transform func(K1, V) K2) R {
	if len(vals) == 0 {
		return nil
	}

	out := make(R, len(vals))
	for key, val := range vals {
		k := transform(key, val)
		out[k] = val
	}
	return out
}
