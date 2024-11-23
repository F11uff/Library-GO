package copy

import (
	"golang.org/x/exp/constraints"
)

func Copy[T constraints.Ordered](to []T, from []T) {
	for i := range to {
		to[i] = from[i]
	}
}
