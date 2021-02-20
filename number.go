package customsortgo

import "sort"

func SortInts(nums []int) []int {
	sort.Sort(sort.IntSlice(nums))
	return nums
}
