package maps

// ContainsKey returns true if target key is found in the map.
func ContainsKey[M ~map[K]V, K comparable, V any](vals M, target K) bool {
	for key := range vals {
		if key == target {
			return true
		}
	}
	return false
}
