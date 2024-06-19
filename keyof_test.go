package maps_test

import (
	"testing"

	"github.com/dosadczuk/go-maps"
	"github.com/google/go-cmp/cmp"
)

func TestKeyOf(t *testing.T) {
	tt := map[string]struct {
		// input
		values map[int]int
		target int
		// assert
		want int
	}{
		"empty map": {
			values: nil, // zero value
			target: 0,
			want:   0, // zero value
		},
		"map with key-value pair": {
			values: map[int]int{1: 2},
			target: 2,
			want:   1,
		},
		"map with key-value pairs": {
			values: map[int]int{1: 2, 2: 1},
			target: 1,
			want:   2,
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			have := maps.KeyOf(tc.values, tc.target)

			if !cmp.Equal(tc.want, have) {
				t.Error(cmp.Diff(tc.want, have))
			}
		})
	}
}
