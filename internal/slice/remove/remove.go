package remove

import (
	"errors"
	"golang.org/x/exp/constraints"
)

func Remove[T constraints.Ordered](list []T, index int) ([]T, error) {
	if index > len(list) {
		return list, errors.New("error : Index out of range")
	}

	var newList []T
	newList = make([]T, len(list)-1, cap(list))

	for i, j := 0, 0; i < len(list); i, j = i+1, j+1 {
		if i == index {
			j--
			continue
		}
		newList[j] = list[i]
	}

	return newList, nil
}
