package pure

import (
	"cmp"
	"sort"
)

type Item interface {
	cmp.Ordered
}

func copySlice[I Item](original []I) []I {
	c := make([]I, len(original))
	copy(c, original)
	return c
}

func NewSorted[I Item](items []I) []I {
	n := copySlice(items)
	sort.Slice(n, func(x, y int) bool {
		return n[x] < n[y]
	})
	return n
}

func NewReverseSorted[I Item](items []I) []I {
	n := copySlice(items)
	sort.Slice(n, func(x, y int) bool {
		return n[x] > n[y]
	})
	return n
}
