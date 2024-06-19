package maps

// ForEach calls the given action on each key-value pair.
func ForEach[M ~map[K]V, K comparable, V any](vals M, action func(K, V)) {
	for key, val := range vals {
		action(key, val)
	}
}
