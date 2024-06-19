package maps_test

import (
	"testing"

	"github.com/dosadczuk/go-maps"
	"github.com/google/go-cmp/cmp"
)

func TestFind(t *testing.T) {
	tt := map[string]struct {
		// input
		values    map[int]int
		predicate func(int, int) bool
		// assert
		wantValue int
		wantFound bool
	}{
		"empty map": {
			values:    nil, // zero value
			predicate: func(_, _ int) bool { return true },
			wantValue: 0,     // zero value
			wantFound: false, // zero value
		},
		"map with value matching predicate": {
			values:    map[int]int{1: 2},
			predicate: func(_, val int) bool { return val%2 == 0 },
			wantValue: 2,
			wantFound: true,
		},
		"map with value not matching predicate": {
			values:    map[int]int{1: 2, 2: 1},
			predicate: func(key, val int) bool { return key > 2 || val > 2 },
			wantValue: 0,
			wantFound: false,
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			haveValue, haveFound := maps.Find(tc.values, tc.predicate)

			if !cmp.Equal(tc.wantValue, haveValue) {
				t.Error(cmp.Diff(tc.wantValue, haveValue))
			}
			if !cmp.Equal(tc.wantFound, haveFound) {
				t.Error(cmp.Diff(tc.wantFound, haveFound))
			}
		})
	}
}
