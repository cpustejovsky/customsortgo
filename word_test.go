package customsortgo_test

import (
	"testing"

	csgo "github.com/cpustejovsky/customsortgo"
	th "github.com/cpustejovsky/customsortgo/testhelp"
)

var testChars csgo.Chars = []rune("hello, world!")

func TestChars(t *testing.T) {
	wantLen := 13
	gotLen := testChars.Len()
	th.AssertEqual(t, gotLen, wantLen)

	wantLessFirst := false
	gotLessFirst := testChars.Less(0, 1)
	th.AssertEqual(t, gotLessFirst, wantLessFirst)

	wantSwap := "ehllo, world!"
	testChars.Swap(0, 1)
	th.AssertEqual(t, string(testChars), wantSwap)

	wantLessSecond := true
	gotLessSecond := testChars.Less(0, 1)
	th.AssertEqual(t, gotLessSecond, wantLessSecond)
}

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
