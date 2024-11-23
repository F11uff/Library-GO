package append

import (
	"golang.org/x/exp/constraints"
)

func Append[T constraints.Ordered](list []T, el ...T) []T {
	var newList []T
	newLen := len(list) + len(el)
	lenList := len(list)

	if newLen <= cap(list) {
		newList = make([]T, newLen, cap(list))
		copy(newList, list)
	} else {
		newCap := newLen
		if newCap < 2*newLen {
			newCap = 2 * newLen
		}

		newList = make([]T, newLen, newCap)
		copy(newList, list)

	}

	for i := 0; i < len(el); i++ {
		newList[lenList] = el[i]
		lenList++
	}

	return newList
}
