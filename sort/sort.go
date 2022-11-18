package sort

import (
	"sort"
)

func NewSortedIntegers(nums []int) []int {
	n := make([]int, len(nums))
	copy(n, nums)
	sort.IntSlice(n).Sort()
	return n
}

func NewSortedFloats(nums []float64) []float64 {
	n := make([]float64, len(nums))
	copy(n, nums)
	sort.Float64Slice(n).Sort()
	return n
}

func NewSortedStrings(strs []string) []string {
	s := make([]string, len(strs))
	copy(s, strs)
	sort.StringSlice(s).Sort()
	return s
}

func NewReverseSortedIntegers(nums []int) []int {
	n := make([]int, len(nums))
	copy(n, nums)
	sort.Sort(sort.Reverse(sort.IntSlice(n)))
	return n
}

func NewReverseSortedFloats(nums []float64) []float64 {
	n := make([]float64, len(nums))
	copy(n, nums)
	sort.Sort(sort.Reverse(sort.Float64Slice(n)))
	return n
}

func NewReverseSortedStrings(strs []string) []string {
	s := make([]string, len(strs))
	copy(s, strs)
	sort.Sort(sort.Reverse(sort.StringSlice(s)))
	return s
}
