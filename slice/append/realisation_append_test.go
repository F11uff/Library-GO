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
			expectedCap: 6,
		},
		{
			name:        "Append to non-empty slice",
			input:       []int{1, 2},
			elements:    []int{3, 4},
			expected:    []int{1, 2, 3, 4},
			expectedLen: 4,
			expectedCap: 8,
		},
		{
			name:        "Append with capacity expansion",
			input:       []int{1, 2, 3, 4},
			elements:    []int{5, 6},
			expected:    []int{1, 2, 3, 4, 5, 6},
			expectedLen: 6,
			expectedCap: 12,
		},
		{
			name:        "Append single element",
			input:       []int{1},
			elements:    []int{2},
			expected:    []int{1, 2},
			expectedLen: 2,
			expectedCap: 4,
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
			expectedCap: 20,
		},
		{
			name:        "Append many elements",
			input:       make([]int, 1000),
			elements:    make([]int, 1000),
			expected:    append(make([]int, 1000), make([]int, 1000)...),
			expectedLen: 2000,
			expectedCap: 4000,
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
}
