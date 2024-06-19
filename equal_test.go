package maps_test

import (
	"testing"

	"github.com/dosadczuk/go-maps"
	"github.com/google/go-cmp/cmp"
)

func TestEqual(t *testing.T) {
	tt := map[string]struct {
		// input
		values1 map[int]int
		values2 map[int]int
		// assert
		want bool
	}{
		"empty maps": {
			values1: nil, // zero value
			values2: nil, // zero value
			want:    true,
		},
		"maps with different size": {
			values1: map[int]int{1: 2},
			values2: map[int]int{1: 2, 2: 3},
			want:    false,
		},
		"maps with the same keys but different values": {
			values1: map[int]int{1: 2},
			values2: map[int]int{1: 3},
			want:    false,
		},
		"maps with the same values but different keys": {
			values1: map[int]int{1: 2},
			values2: map[int]int{2: 2},
			want:    false,
		},
		"maps with the same key-value pairs": {
			values1: map[int]int{1: 2, 2: 3},
			values2: map[int]int{1: 2, 2: 3},
			want:    true,
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			have := maps.Equal(tc.values1, tc.values2)

			if !cmp.Equal(tc.want, have) {
				t.Error(cmp.Diff(tc.want, have))
			}
		})
	}
}

func TestEqualFunc(t *testing.T) {
	tt := map[string]struct {
		// input
		values1 map[int]int
		values2 map[int]int
		// assert
		want bool
	}{
		"empty maps": {
			values1: nil, // zero value
			values2: nil, // zero value
			want:    true,
		},
		"maps with different size": {
			values1: map[int]int{1: 2},
			values2: map[int]int{1: 2, 2: 3},
			want:    false,
		},
		"maps with the same keys but different values": {
			values1: map[int]int{1: 2},
			values2: map[int]int{1: 3},
			want:    false,
		},
		"maps with the same values but different keys": {
			values1: map[int]int{1: 2},
			values2: map[int]int{2: 2},
			want:    false,
		},
		"maps with the same key-value pairs": {
			values1: map[int]int{1: 2, 2: 3},
			values2: map[int]int{1: 2, 2: 3},
			want:    true,
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			have := maps.EqualFunc(tc.values1, tc.values2, func(v1, v2 int) bool { return v1 == v2 })

			if !cmp.Equal(tc.want, have) {
				t.Error(cmp.Diff(tc.want, have))
			}
		})
	}
}
