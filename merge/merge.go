package merge

func Sort(nums []int, left, right int) {
	if left < right {
		//calculate the middle by taking the average of the difference between right and left
		//for a full array it would be len(arr) - 0 /2 or len(arr)/2
		//as you break down the array the starting index (left) won't be zero so you need to add that to the middle to correctly offset it
		middle := left + (right-left)/2
		//break this array down further from the left to the middle of the array
		Sort(nums, left, middle)
		//break this array down further from the middle to the right of the array
		Sort(nums, middle+1, right)
		merge(nums, left, right, middle)
	}
}

func merge(array []int, left, right, mid int) {
	leftArr := make([]int, mid-left+1)
	rightArr := make([]int, right-mid)

	copy(leftArr, array[left:mid+1])
	copy(rightArr, array[mid+1:right+1])

	var indexOfLeftArr, indexOfRightArr int
	indexOfMergedArr := left

	for indexOfLeftArr < len(leftArr) && indexOfRightArr < len(rightArr) {
		if leftArr[indexOfLeftArr] <= rightArr[indexOfRightArr] {
			array[indexOfMergedArr] = leftArr[indexOfLeftArr]
			indexOfLeftArr++
		} else {
			array[indexOfMergedArr] = rightArr[indexOfRightArr]
			indexOfRightArr++
		}
		indexOfMergedArr++
	}

	for indexOfLeftArr < len(leftArr) {
		array[indexOfMergedArr] = leftArr[indexOfLeftArr]
		indexOfLeftArr++
		indexOfMergedArr++
	}
	for indexOfRightArr < len(rightArr) {
		array[indexOfMergedArr] = rightArr[indexOfRightArr]
		indexOfRightArr++
		indexOfMergedArr++
	}
}

func NewSorted(nums []int) []int {
	var n []int
	copy(n, nums)
	Sort(n, 0, len(n)-1)
	return n
}
