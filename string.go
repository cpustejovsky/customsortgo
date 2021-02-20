package string

import (
	"sort"
)

//Words type is what a list of strings are converted into to act as a sort Interface
type Words []string

func (w Words) Len() int {
	return len(w)
}
func (w Words) Swap(i, j int) {
	w[i], w[j] = w[j], w[i]
}
func (w Words) Less(i, j int) bool {
	return w[i] < w[j]
}

//SortWords takes a slice of strings and returns them in a sorted order
func SortWords(words []string) []string {
	w := Words(words)
	sort.Sort(w)
	return []string(w)
}


//ReverseSortWords takes a slice of strings and returns them in a reverse sorted order
func ReverseSortWords(words []string) []string {
	w := Words((words))
	sort.Sort(sort.Reverse(w))
	return []string(w)
}
