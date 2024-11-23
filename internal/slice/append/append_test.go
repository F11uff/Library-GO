package main

import (
	"testing"
)

type AppendTestCase[T any] struct {
	name        string
	input       []T
	elements    []T
	expected    []T
	expectedLen int
	expectedCap int
}

func TestAppend(t *testing.T) {
	tests := []AppendTestCase[int]{
		{
			name:        "Append to empty slice",
			input:       []int{},
			elements:    []int{1, 2, 3},
			expected:    []int{1, 2, 3},
			expectedLen: 3,
			expectedCap: 3,
		},
		{
			name:        "Append to non-empty slice",
			input:       []int{1, 2},
			elements:    []int{3, 4},
			expected:    []int{1, 2, 3, 4},
			expectedLen: 4,
			expectedCap: 4,
		},
		{
			name:        "Append with capacity expansion",
			input:       []int{1, 2, 3, 4},
			elements:    []int{5, 6},
			expected:    []int{1, 2, 3, 4, 5, 6},
			expectedLen: 6,
			expectedCap: 8,
		},
		{
			name:        "Append single element",
			input:       []int{1},
			elements:    []int{2},
			expected:    []int{1, 2},
			expectedLen: 2,
			expectedCap: 2,
		},
		{
			name:        "Append empty elements",
			input:       []int{1, 2, 3},
			elements:    []int{},
			expected:    []int{1, 2, 3},
			expectedLen: 3,
			expectedCap: 3,
		},
		{
			name:        "Append zero elements to empty slice",
			input:       []int{},
			elements:    []int{},
			expected:    []int{},
			expectedLen: 0,
			expectedCap: 0,
		},
		{
			name:        "Append with doubled capacity",
			input:       []int{1, 2, 3, 4, 5, 6, 7, 8},
			elements:    []int{9, 10},
			expected:    []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			expectedLen: 10,
			expectedCap: 16,
		},
		{
			name:        "Append many elements",
			input:       make([]int, 1000),
			elements:    make([]int, 1000),
			expected:    append(make([]int, 1000), make([]int, 1000)...),
			expectedLen: 2000,
			expectedCap: 2000,
		},
	}

	stringTests := []AppendTestCase[string]{
		{
			name:        "Append to empty slice of strings",
			input:       []string{},
			elements:    []string{"a", "b", "c"},
			expected:    []string{"a", "b", "c"},
			expectedLen: 3,
			expectedCap: 4,
		},
		{
			name:        "Append to non-empty slice of strings",
			input:       []string{"hello", "world"},
			elements:    []string{"golang", "rocks"},
			expected:    []string{"hello", "world", "golang", "rocks"},
			expectedLen: 4,
			expectedCap: 4,
		},
		{
			name:        "Append empty elements to slice of strings",
			input:       []string{"apple", "banana"},
			elements:    []string{},
			expected:    []string{"apple", "banana"},
			expectedLen: 2,
			expectedCap: 2,
		},
	}

	floatTests := []AppendTestCase[float64]{
		{
			name:        "Append to empty slice of floats",
			input:       []float64{},
			elements:    []float64{1.1, 2.2, 3.3},
			expected:    []float64{1.1, 2.2, 3.3},
			expectedLen: 3,
			expectedCap: 4,
		},
		{
			name:        "Append to non-empty slice of floats",
			input:       []float64{1.1, 2.2},
			elements:    []float64{3.3, 4.4},
			expected:    []float64{1.1, 2.2, 3.3, 4.4},
			expectedLen: 4,
			expectedCap: 4,
		},
		{
			name:        "Append with capacity expansion for floats",
			input:       []float64{1.1, 2.2, 3.3, 4.4},
			elements:    []float64{5.5, 6.6},
			expected:    []float64{1.1, 2.2, 3.3, 4.4, 5.5, 6.6},
			expectedLen: 6,
			expectedCap: 8,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Append(tt.input, tt.elements...)

			if len(result) != tt.expectedLen {
				t.Errorf("%s: expected length %d, got %d", tt.name, tt.expectedLen, len(result))
			}

			if cap(result) < tt.expectedCap {
				t.Errorf("%s: expected capacity at least %d, got %d", tt.name, tt.expectedCap, cap(result))
			}

			for i := range tt.expected {
				if result[i] != tt.expected[i] {
					t.Errorf("%s: expected element at index %d to be %v, got %v", tt.name, i, tt.expected[i], result[i])
				}
			}
		})
	}

	for _, tt := range stringTests {
		t.Run(tt.name, func(t *testing.T) {
			result := Append(tt.input, tt.elements...)
			if len(result) != tt.expectedLen {
				t.Errorf("%s: expected length %d, got %d", tt.name, tt.expectedLen, len(result))
			}
			if cap(result) < tt.expectedCap {
				t.Errorf("%s: expected capacity at least %d, got %d", tt.name, tt.expectedCap, cap(result))
			}
			for i := range tt.expected {
				if result[i] != tt.expected[i] {
					t.Errorf("%s: expected element at index %d to be %v, got %v", tt.name, i, tt.expected[i], result[i])
				}
			}
		})
	}

	for _, tt := range floatTests {
		t.Run(tt.name, func(t *testing.T) {
			result := Append(tt.input, tt.elements...)
			if len(result) != tt.expectedLen {
				t.Errorf("%s: expected length %d, got %d", tt.name, tt.expectedLen, len(result))
			}
			if cap(result) < tt.expectedCap {
				t.Errorf("%s: expected capacity at least %d, got %d", tt.name, tt.expectedCap, cap(result))
			}
			for i := range tt.expected {
				if result[i] != tt.expected[i] {
					t.Errorf("%s: expected element at index %d to be %v, got %v", tt.name, i, tt.expected[i], result[i])
				}
			}
		})
	}
}
