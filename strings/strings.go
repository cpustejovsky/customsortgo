package strings

import (
	"sort"
)

// Strings type is a slice of Strings
type Strings []string

func (s Strings) Len() int {
	return len(s)
}
func (s Strings) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s Strings) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}

func SortByLength(s []string) {
	sort.Sort(Strings(s))
}
