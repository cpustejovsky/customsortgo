package customsortgo

import (
	"sort"
)

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

//SortWord takes a string and returns it with it's characters in a sorted order
func SortWord(word string) string {
	ch := Chars(word)
	sort.Sort(ch)
	return string(ch)
}

//ReverseSortWord takes a string and returns it with it's characters in a reverse sorted order
func ReverseSortWord(word string) string {
	ch := Chars(word)
	sort.Sort(sort.Reverse(ch))
	return string(ch)
}
