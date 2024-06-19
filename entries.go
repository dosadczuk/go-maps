package maps

// Entry is just a key-value pair.
type Entry[K comparable, V any] struct {
	Key   K
	Value V
}

// Entries returns key-value pairs from the map, order is random.
func Entries[M ~map[K]V, R []Entry[K, V], K comparable, V any](vals M) R {
	if len(vals) == 0 {
		return nil
	}

	out := make(R, 0, len(vals))
	for key, val := range vals {
		out = append(out, Entry[K, V]{key, val})
	}
	return out
}
