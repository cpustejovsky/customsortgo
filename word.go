package customsortgo

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

//Chars type is what a string is converted into to act as a sort Interface
type Chars []rune

func (ch Chars) Len() int {
	return len(ch)
}
func (ch Chars) Swap(i, j int) {
	ch[i], ch[j] = ch[j], ch[i]
}
func (ch Chars) Less(i, j int) bool {
	return ch[i] < ch[j]
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

//SortWord takes a string and returns it with it's characters in a sorted order
func SortWord(word string) string {
	var sw Chars
	for _, c := range word {
		sw = append(sw, c)
	}
	sort.Sort(sw)
	return string(sw)
}

//ReverseSortWord takes a string and returns it with it's characters in a reverse sorted order
func ReverseSortWord(word string) string {
	var sw Chars
	for _, c := range word {
		sw = append(sw, c)
	}
	sort.Sort(sort.Reverse(sw))
	return string(sw)
}