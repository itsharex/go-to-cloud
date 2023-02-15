package utils

import (
	"reflect"
	"testing"
)

func TestIntersect(t *testing.T) {
	testCases := []struct {
		a, b, want []uint
	}{
		{[]uint{1, 2, 3}, []uint{2, 3, 4}, []uint{2, 3}},
		{[]uint{1, 2, 3}, []uint{4, 5, 6}, []uint{}},
		{[]uint{1, 2, 3}, []uint{}, []uint{}},
		{[]uint{}, []uint{2, 3, 4}, []uint{}},
		{[]uint{}, []uint{}, []uint{}},
	}
	for _, tc := range testCases {
		if got := Intersect(tc.a, tc.b); !reflect.DeepEqual(got, tc.want) {
			t.Errorf("Intersect(%v, %v) = %v, want %v", tc.a, tc.b, got, tc.want)
		}
	}
}
