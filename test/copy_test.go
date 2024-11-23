package test

import (
	c "Library/internal/slice/copy"
	"errors"
	"golang.org/x/exp/constraints"
	"testing"
)

type Checker interface {
	Check(t *testing.T)
}

type TestCaseCopy[T constraints.Ordered] struct {
	CaseNameTest string
	List1        []T
	List2        []T
	OutputLen    int64
	OutputCap    int64
	OutputList   []T
}

func realisationCheck(checker Checker, t *testing.T) {
	checker.Check(t)
}

func (test *TestCaseCopy[T]) Check(t *testing.T) {
	t.Run(test.CaseNameTest, func(t *testing.T) {
		c.Copy(test.List1, test.List2)
		err := CheckSlice(test.List1, test.OutputList)
		if err != nil {
			t.Errorf("Error : test fail - %s, List1 - %v, OutputList - %v", test.CaseNameTest, test.List1, test.OutputList)
		} else if test.OutputLen != int64(len(test.List1)) {
			t.Errorf("Error : len List1 - %v, OutputList - %v", len(test.List1), test.OutputLen)
		} else if test.OutputCap != test.OutputCap {
			t.Errorf(" Error : cap List1 - %v, OutputList - %v", cap(test.List1), test.OutputCap)
		}
	})
}

func CheckSlice[T comparable](list1, list2 []T) error {
	for i := range list1 {
		if list1[i] != list2[i] {
			return errors.New("not equal")
		}
	}

	return nil
}

func TestCopy(t *testing.T) {
	testsCaseInt := []TestCaseCopy[int]{
		{
			CaseNameTest: "1 (int)",
			List1:        []int{0, 0, 0, 0, 0},
			List2:        []int{1, 1, 1, 1, 1},
			OutputLen:    5,
			OutputCap:    5,
			OutputList:   []int{1, 1, 1, 1, 1},
		},
		{
			CaseNameTest: "2 (int)",
			List1:        make([]int, 5, 5),
			List2:        []int{1, 5, 2, 7, 9},
			OutputLen:    5,
			OutputCap:    5,
			OutputList:   []int{1, 5, 2, 7, 9},
		},
		{
			CaseNameTest: "3 (int)",
			List1:        make([]int, 1, 1),
			List2:        make([]int, 1, 1),
			OutputLen:    1,
			OutputCap:    1,
			OutputList:   []int{0},
		},
	}

	testsCaseFloat := []TestCaseCopy[float64]{
		{
			CaseNameTest: "1 (float)",
			List1:        []float64{0.0, 0.0, 0.0, 0.0, 0.0},
			List2:        []float64{1.1, 1.1, 1.1, 1.1, 1.1},
			OutputLen:    5,
			OutputCap:    5,
			OutputList:   []float64{1.1, 1.1, 1.1, 1.1, 1.1},
		},
		{
			CaseNameTest: "2 (float)",
			List1:        make([]float64, 3, 6),
			List2:        []float64{4.5, 123.4, 76.0},
			OutputLen:    3,
			OutputCap:    6,
			OutputList:   []float64{4.5, 123.4, 76.0},
		},
		{
			CaseNameTest: "3 (float)",
			List1:        make([]float64, 1, 1),
			List2:        make([]float64, 1, 1),
			OutputLen:    1,
			OutputCap:    1,
			OutputList:   []float64{0},
		},
	}

	testsCaseString := []TestCaseCopy[string]{
		{
			CaseNameTest: "1 (string)",
			List1:        []string{"a", "b", "c"},
			List2:        []string{"a", "b", "c"},
			OutputLen:    3,
			OutputCap:    3,
			OutputList:   []string{"a", "b", "c"},
		},
		{
			CaseNameTest: "2 (string)",
			List1:        []string{"a", "b", "d"},
			List2:        []string{"a", "b", "c"},
			OutputLen:    3,
			OutputCap:    3,
			OutputList:   []string{"a", "b", "c"},
		},
		{
			CaseNameTest: "3 (string)",
			List1:        make([]string, 1, 1),
			List2:        []string{"a"},
			OutputLen:    1,
			OutputCap:    1,
			OutputList:   []string{"a"},
		},
	}

	for _, test := range testsCaseInt {
		realisationCheck(&test, t)
	}

	for _, test := range testsCaseFloat {
		realisationCheck(&test, t)
	}

	for _, test := range testsCaseString {
		realisationCheck(&test, t)
	}
}
