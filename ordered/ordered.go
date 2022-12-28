package ordered

import (
	"github.com/gokch/goslice"
	"golang.org/x/exp/constraints"
	"golang.org/x/exp/slices"
)

type Slice[T constraints.Ordered] struct {
	goslice.Slice[T]
}

func (s *Slice[T]) Sort() {
	slices.Sort(s.Slice)
}

func (s *Slice[T]) SortFunc(less func(a, b T) bool) {
	slices.SortFunc(s.Slice, less)
}

func (s *Slice[T]) IsSorted() bool {
	return slices.IsSorted(s.Slice)
}

func (s *Slice[T]) Search(x T) (int, bool) {
	return slices.BinarySearch(s.Slice, x)
}

func (s *Slice[T]) SearchFunc(x T, cmp func(a, b T) int) (int, bool) {
	return slices.BinarySearchFunc(s.Slice, x, cmp)
}
