package maps_test

import (
	"testing"

	"github.com/dosadczuk/go-maps"
	"github.com/google/go-cmp/cmp"
)

func TestContainsValue(t *testing.T) {
	tt := map[string]struct {
		// input
		values map[int]int
		target int
		// assert
		want bool
	}{
		"empty map": {
			values: nil, // zero value
			target: 1,
			want:   false,
		},
		"map with matching value": {
			values: map[int]int{1: 2},
			target: 2,
			want:   true,
		},
		"map with not matching value": {
			values: map[int]int{1: 2},
			target: 1,
			want:   false,
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			have := maps.ContainsValue(tc.values, tc.target)

			if !cmp.Equal(tc.want, have) {
				t.Error(cmp.Diff(tc.want, have))
			}
		})
	}
}
