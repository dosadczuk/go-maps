package maps_test

import (
	"testing"

	"github.com/dosadczuk/go-maps"
	"github.com/google/go-cmp/cmp"
)

func TestMapKeys(t *testing.T) {
	tt := map[string]struct {
		// input
		values    map[int]int
		transform func(int, int) int
		// assert
		want map[int]int
	}{
		"empty map": {
			values:    nil, // zero value
			transform: func(_, _ int) int { return 0 },
			want:      nil, // zero value
		},
		"map with key-value pair": {
			values:    map[int]int{1: 2},
			transform: func(key, val int) int { return key + val },
			want:      map[int]int{3: 2},
		},
		"map with key-value pairs": {
			values: map[int]int{
				1: 2,
				2: 4,
				3: 6,
			},
			transform: func(key, val int) int { return key + val },
			want: map[int]int{
				3: 2,
				6: 4,
				9: 6,
			},
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			have := maps.MapKeys(tc.values, tc.transform)

			if !cmp.Equal(tc.want, have) {
				t.Error(cmp.Diff(tc.want, have))
			}
		})
	}
}
