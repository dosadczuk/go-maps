package maps_test

import (
	"sort"
	"testing"

	"github.com/dosadczuk/go-maps"
	"github.com/google/go-cmp/cmp"
)

func TestKeys(t *testing.T) {
	tt := map[string]struct {
		// input
		values map[int]int
		// assert
		want []int
	}{
		"empty map": {
			values: nil, // zero value
			want:   nil, // zero value
		},
		"map with key-value pair": {
			values: map[int]int{1: 2},
			want:   []int{1},
		},
		"map with key-value pairs": {
			values: map[int]int{1: 2, 2: 1},
			want:   []int{1, 2},
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			have := maps.Keys(tc.values)
			sort.Ints(have)

			if !cmp.Equal(tc.want, have) {
				t.Error(cmp.Diff(tc.want, have))
			}
		})
	}
}
