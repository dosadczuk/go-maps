package maps

// Max returns the largest value from the map.
func Max[M ~map[K]V, K comparable, V orderable](vals M) (max V) {
	if len(vals) == 0 {
		return
	}

	// hack for not knowing the first key-value pair
	for _, val := range vals {
		max = val
		break
	}
	for _, val := range vals {
		if val > max {
			max = val
		}
	}
	return
}

// MaxFunc returns the largest value from the map.
//
// Function compares two arguments for order and returns:
//
//	 0 - if the arguments are equal,
//	-1 - if the first argument is less than the second,
//	+1 - if the second argument is less than the first.
func MaxFunc[M ~map[K]V, K comparable, V any](vals M, cmp func(V, V) int) (max V) {
	if len(vals) == 0 {
		return
	}

	// hack for not knowing the first key-value pair
	for _, val := range vals {
		max = val
		break
	}
	for _, val := range vals {
		if cmp(val, max) > 0 {
			max = val
		}
	}
	return
}
