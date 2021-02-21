package customsortgo_test

import (
	"testing"

	csgo "github.com/cpustejovsky/customsortgo"
	th "github.com/cpustejovsky/customsortgo/testhelp"
)

func TestSortWord(t *testing.T) {
	w := "cbadegf"
	want := "abcdefg"
	got := csgo.SortWord(w)

	th.AssertEqual(t, got, want)

}

func TestReverseSortWord(t *testing.T) {
	w := "cbadegf"
	want := "gfedcba"
	got := csgo.ReverseSortWord(w)

	th.AssertEqual(t, got, want)

}