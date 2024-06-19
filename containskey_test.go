package maps_test

import (
	"testing"

	"github.com/dosadczuk/go-maps"
	"github.com/google/go-cmp/cmp"
)

func TestContainsKey(t *testing.T) {
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
		"map with matching key": {
			values: map[int]int{1: 2},
			target: 1,
			want:   true,
		},
		"map with not matching key": {
			values: map[int]int{1: 2},
			target: 2,
			want:   false,
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			have := maps.ContainsKey(tc.values, tc.target)

			if !cmp.Equal(tc.want, have) {
				t.Error(cmp.Diff(tc.want, have))
			}
		})
	}
}
