package maps_test

import (
	"sort"
	"testing"

	"github.com/dosadczuk/go-maps"
	"github.com/google/go-cmp/cmp"
)

func TestEntries(t *testing.T) {
	tt := map[string]struct {
		// input
		values map[int]int
		// assert
		want []maps.Entry[int, int]
	}{
		"empty map": {
			values: nil, // zero value
			want:   nil, // zero value
		},
		"map with key-value pair": {
			values: map[int]int{1: 2},
			want:   []maps.Entry[int, int]{{1, 2}},
		},
		"map with key-value pairs": {
			values: map[int]int{
				1: 2,
				2: 3,
				3: 4,
			},
			want: []maps.Entry[int, int]{
				{1, 2},
				{2, 3},
				{3, 4},
			},
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			have := maps.Entries(tc.values)
			sort.Slice(have, func(i, j int) bool {
				return have[i].Key < have[j].Key
			})

			if !cmp.Equal(tc.want, have) {
				t.Error(cmp.Diff(tc.want, have))
			}
		})
	}
}
