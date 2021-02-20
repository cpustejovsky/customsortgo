package customsortgo

import "sort"

type SortIntList []int

func (n SortIntList) Len() int {
	return len(n)
}
func (n SortIntList) Swap(i, j int) {
	n[i], n[j] = n[j], n[i]
}
func (n SortIntList) Less(i, j int) bool {
	return n[i] < n[j]
}


func SortInts(nums []int) []int {
	sort.Sort(SortIntList(nums))
	return nums
}
