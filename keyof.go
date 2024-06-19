package maps

// KeyOf returns key of the first value matching the target.
// Return zero value if not in the map.
func KeyOf[M ~map[K]V, K, V comparable](vals M, target V) (key K) {
	for key, val := range vals {
		if val == target {
			return key
		}
	}
	return
}
