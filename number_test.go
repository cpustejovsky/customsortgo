package customsortgo_test

import (
	"testing"

	csgo "github.com/cpustejovsky/customsortgo"
	th "github.com/cpustejovsky/customsortgo/testhelp"
)

func TestSortInts(t *testing.T) {
	n := []int{1, 3, 5, 2, 7, 6, 4}
	want := []int{1, 2, 3, 4, 5, 6, 7}
	got := csgo.SortInts(n)

	th.AssertEqual(t, got, want)

}
