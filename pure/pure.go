package pure

import (
	"sort"
)

type Item interface {
	int | float64 | string
}

func reserveSort(x sort.Interface) {
	sort.Sort(sort.Reverse(x))
}

func copySlice[I Item](original []I) []I {
	c := make([]I, len(original))
	copy(c, original)
	return c
}

func NewSortedIntegers(ints []int) []int {
	n := copySlice(ints)
	sort.IntSlice(n).Sort()
	return n
}

func NewSortedFloats(float64s []float64) []float64 {
	n := copySlice(float64s)
	sort.Float64Slice(n).Sort()
	return n
}

func NewSortedStrings(strings []string) []string {
	s := copySlice(strings)
	sort.StringSlice(s).Sort()
	return append(s)
}

func NewReverseSortedIntegers(integers []int) []int {
	n := copySlice(integers)
	reserveSort(sort.IntSlice(n))
	return n
}

func NewReverseSortedFloats(float64s []float64) []float64 {
	n := copySlice(float64s)
	reserveSort(sort.Float64Slice(n))
	return n
}

func NewReverseSortedStrings(strings []string) []string {
	s := copySlice(strings)
	reserveSort(sort.StringSlice(s))
	return s
}
