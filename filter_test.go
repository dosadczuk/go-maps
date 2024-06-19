package maps_test

import (
	"testing"

	"github.com/dosadczuk/go-maps"
	"github.com/google/go-cmp/cmp"
)

func TestFilter(t *testing.T) {
	tt := map[string]struct {
		// input
		values    map[int]int
		predicate func(int, int) bool
		// assert
		want map[int]int
	}{
		"empty map": {
			values:    nil, // zero value
			predicate: func(_ int, _ int) bool { return true },
			want:      nil, // zero value
		},
		"map with key-value pair": {
			values:    map[int]int{1: 2},
			predicate: func(_ int, val int) bool { return val%2 == 0 },
			want:      map[int]int{1: 2},
		},
		"map with key-value pairs": {
			values: map[int]int{
				1: 2,
				2: 4,
				3: 6,
			},
			predicate: func(key int, _ int) bool { return key%2 != 0 },
			want: map[int]int{
				1: 2,
				3: 6,
			},
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			have := maps.Filter(tc.values, tc.predicate)

			if !cmp.Equal(tc.want, have) {
				t.Error(cmp.Diff(tc.want, have))
			}
		})
	}
}
