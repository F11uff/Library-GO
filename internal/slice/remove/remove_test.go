package main

import (
	"testing"
)

type TestCase[T any] struct {
	Name     string
	List     []T
	Index    int
	Expected []T
}

func TestRemoveInt(t *testing.T) {
	tests := []TestCase[int]{
		{"Remove from int slice", []int{1, 2, 3, 4, 5}, 2, []int{1, 2, 4, 5}},
		{"Remove first element", []int{1, 2, 3}, 0, []int{2, 3}},
		{"Remove last element", []int{1, 2, 3}, 2, []int{1, 2}},
	}

	for _, tt := range tests {
		result, err := Remove(tt.List, tt.Index)
		if err != nil {
			t.Errorf("%s: unexpected error: %v", tt.Name, err)
			continue
		}
		if !equal(result, tt.Expected) {
			t.Errorf("%s: expected %v, got %v", tt.Name, tt.Expected, result)
		}
	}

	_, err := Remove([]int{1, 2, 3}, 5)
	if err == nil || err.Error() != "error : Index out of range" {
		t.Error("expected index out of range error, got:", err)
	}
}

func TestRemoveFloat64(t *testing.T) {
	tests := []TestCase[float64]{
		{"Remove from float64 slice", []float64{1.1, 2.2, 3.3, 4.4, 5.5}, 2, []float64{1.1, 2.2, 4.4, 5.5}},
		{"Remove first element", []float64{1.1, 2.2, 3.3}, 0, []float64{2.2, 3.3}},
		{"Remove last element", []float64{1.1, 2.2, 3.3}, 2, []float64{1.1, 2.2}},
	}

	for _, tt := range tests {
		result, err := Remove(tt.List, tt.Index)
		if err != nil {
			t.Errorf("%s: unexpected error: %v", tt.Name, err)
			continue
		}
		if !equal(result, tt.Expected) {
			t.Errorf("%s: expected %v, got %v", tt.Name, tt.Expected, result)
		}
	}

	_, err := Remove([]float64{1.1, 2.2, 3.3}, 5)
	if err == nil || err.Error() != "error : Index out of range" {
		t.Error("expected index out of range error, got:", err)
	}
}

func TestRemoveString(t *testing.T) {
	tests := []TestCase[string]{
		{"Remove from string slice", []string{"a", "b", "c", "d", "e"}, 2, []string{"a", "b", "d", "e"}},
		{"Remove first element", []string{"a", "b", "c"}, 0, []string{"b", "c"}},
		{"Remove last element", []string{"a", "b", "c"}, 2, []string{"a", "b"}},
	}

	for _, tt := range tests {
		result, err := Remove(tt.List, tt.Index)
		if err != nil {
			t.Errorf("%s: unexpected error: %v", tt.Name, err)
			continue
		}
		if !equal(result, tt.Expected) {
			t.Errorf("%s: expected %v, got %v", tt.Name, tt.Expected, result)
		}
	}

	_, err := Remove([]string{"a", "b", "c"}, 5)
	if err == nil || err.Error() != "error : Index out of range" {
		t.Error("expected index out of range error, got:", err)
	}
}

func equal[T comparable](a, b []T) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
