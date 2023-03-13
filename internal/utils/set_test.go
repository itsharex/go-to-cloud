package utils

import (
	"reflect"
	"testing"
)

// TestSets 集合测试
func TestSets(t *testing.T) {
	s := New[int](1, 2, 3)

	if Count(s) != 3 {
		t.Errorf("Expected 3, got %d", Count(s))
	}

	if !Has(s, 1, 2, 3) {
		t.Errorf("Expected true, got false")
	}

	if Has(s, 4) {
		t.Errorf("Expected false, got true")
	}

	Remove(s, 2)
	if Count(s) != 2 {
		t.Errorf("Expected 2, got %d", Count(s))
	}

	if !reflect.DeepEqual(List(s), []int{1, 3}) {
		t.Errorf("Expected [1, 3], got %v", List(s))
	}

	s2 := New[int](3, 4, 5)
	u := Union(s, s2)
	SortList(u)
	su := New[int](1, 3, 4, 5)
	SortList(su)
	if !reflect.DeepEqual(u, su) {
		t.Errorf("Expected [1, 3, 4, 5], got %v", List(u))
	}

	m := Minus(s, s2)
	if !reflect.DeepEqual(List(m), []int{1}) {
		t.Errorf("Expected [1], got %v", List(m))
	}

	c := Complement(s, New[int](1, 2, 3, 4, 5))
	if !reflect.DeepEqual(SortList(c), []int{2, 4, 5}) {
		t.Errorf("Expected [2, 4, 5], got %v", List(c))
	}
}
