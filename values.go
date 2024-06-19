package maps

// Values returns values of the map.
func Values[M ~map[K]V, R []V, K comparable, V any](vals M) R {
	if len(vals) == 0 {
		return nil
	}

	out := make(R, 0, len(vals))
	for _, val := range vals {
		out = append(out, val)
	}
	return out
}
