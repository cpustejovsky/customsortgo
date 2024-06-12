package count

func Sort(nums []int, max int) {
	countSort := make([]int, max+1)
	for _, num := range nums {
		countSort[num]++
	}
	var countSortIndex, i int
	for i < len(nums) {
		for countSort[countSortIndex] > 0 {
			nums[i] = countSortIndex
			countSort[countSortIndex]--
			i++
		}
		countSortIndex++
	}
}

func NewSorted(nums []int, max int) []int {
	n := make([]int, len(nums))
	copy(n, nums)
	Sort(n, max)
	return n
}
