package test

import (
	c "Library/internal/slice/copy"
	"errors"
	"testing"
)

type TestCaseCopy[T any] struct {
	CaseNameTest string
	List1        []T
	List2        []T
	OutputLen    int64
	OutputCap    int64
	OutputList   []T
}

func CheckSlice[T comparable](list1, list2 []T) error {
	for i := range list1 {
		if list1[i] != list2[i] {
			return errors.New("not equal")
		}
	}

	return nil
}

func TestCopyInt(t *testing.T) {
	testsCaseInt := []TestCaseCopy[int]{
		{
			CaseNameTest: "1",
			List1:        []int{0, 0, 0, 0, 0},
			List2:        []int{1, 1, 1, 1, 1},
			OutputLen:    5,
			OutputCap:    5,
			OutputList:   []int{1, 1, 1, 1, 1},
		},
		//{
		//	CaseNameTest: "2",
		//	List1:        make([]int, 5, 5),
		//},
	}

	for _, test := range testsCaseInt {
		c.Copy(test.List1, test.List2)
		err := CheckSlice(test.List1, test.OutputList)
		if err != nil {
			t.Errorf("Error : test fail - %s, List1 - %v, OutputList - %v", test.CaseNameTest, test.List1, test.OutputList)
		} else if test.OutputLen != int64(len(test.List1)) {
			t.Errorf("Error : len List1 - %v, OutputList - %v", len(test.List1), test.OutputLen)
		} else if test.OutputCap != test.OutputCap {
			t.Errorf(" Error : cap List1 - %v, OutputList - %v", cap(test.List1), test.OutputCap)
		}
	}
}
