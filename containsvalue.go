package maps

// ContainsValue returns true if target value is found in the map.
func ContainsValue[M ~map[K]V, K, V comparable](vals M, target V) bool {
	for _, val := range vals {
		if val == target {
			return true
		}
	}
	return false
}
