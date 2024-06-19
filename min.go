package maps

// Min returns the smallest value from the map.
func Min[M ~map[K]V, K comparable, V orderable](vals M) (min V) {
	if len(vals) == 0 {
		return
	}

	// hack for not knowing the first key-value pair
	for _, val := range vals {
		min = val
		break
	}
	for _, val := range vals {
		if val < min {
			min = val
		}
	}
	return
}

// MinFunc returns the smallest value from the map.
//
// Function compares two arguments for order and returns:
//
//	 0 - if the arguments are equal,
//	-1 - if the first argument is less than the second,
//	+1 - if the second argument is less than the first.
func MinFunc[M ~map[K]V, K comparable, V any](vals M, cmp func(V, V) int) (min V) {
	if len(vals) == 0 {
		return
	}

	// hack for not knowing the first key-value pair
	for _, val := range vals {
		min = val
		break
	}
	for _, val := range vals {
		if cmp(val, min) < 0 {
			min = val
		}
	}
	return
}
