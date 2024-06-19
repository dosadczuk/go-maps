package maps_test

import (
	"testing"

	"github.com/dosadczuk/go-maps"
	"github.com/google/go-cmp/cmp"
)

func TestMin(t *testing.T) {
	tt := map[string]struct {
		// input
		values map[int]int
		// assert
		want int
	}{
		"empty map": {
			values: nil, // zero value
			want:   0,   // zero value
		},
		"map with key-value pair": {
			values: map[int]int{1: 2},
			want:   2,
		},
		"map with key-value pairs": {
			values: map[int]int{
				3: 3,
				1: 1,
				2: 4,
			},
			want: 1,
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			have := maps.Min(tc.values)

			if !cmp.Equal(tc.want, have) {
				t.Error(cmp.Diff(tc.want, have))
			}
		})
	}
}

func TestMinFunc(t *testing.T) {
	tt := map[string]struct {
		// input
		values map[int]int
		// assert
		want int
	}{
		"empty map": {
			values: nil, // zero value
			want:   0,   // zero value
		},
		"map with key-value pair": {
			values: map[int]int{1: 2},
			want:   2,
		},
		"map with key-value pairs": {
			values: map[int]int{
				3: 3,
				1: 1,
				2: 4,
			},
			want: 1,
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			have := maps.MinFunc(tc.values, func(v1, v2 int) int { return v1 - v2 })

			if !cmp.Equal(tc.want, have) {
				t.Error(cmp.Diff(tc.want, have))
			}
		})
	}
}
