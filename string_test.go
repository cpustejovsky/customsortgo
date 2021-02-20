package customsortgo_test

import (
	"testing"

	csgo "github.com/cpustejovsky/customsortgo"
	th "github.com/cpustejovsky/customsortgo/testhelp"
)

func TestSortWords(t *testing.T) {
	w := []string{"cat", "albatross", "dolphin", "bee", "zebra", "aardvark"}
	want := []string{"aardvark", "albatross", "bee", "cat", "dolphin", "zebra"}
	got := csgo.SortWords(w)

	th.AssertEqual(t, got, want)

}

func TestReverseSortWords(t *testing.T) {
	w := []string{"cat", "albatross", "dolphin", "bee", "zebra", "aardvark"}
	want := []string{"zebra", "dolphin", "cat", "bee", "albatross", "aardvark"}
	got := csgo.ReverseSortWords(w)

	th.AssertEqual(t, got, want)

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