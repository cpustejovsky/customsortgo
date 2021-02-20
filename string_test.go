package string_test

import (
	"testing"
	th "github.com/cpustejovsky/customsortgo/testhelp"
	csgo "github.com/cpustejovsky/customsortgo"
)

func TestSortWords(t *testing.T){
	w := []string{"cat", "albatross", "dolphin", "bee", "zebra", "aardvark"}
	want := []string{"aardvark", "albatross", "bee", "cat", "dolphin", "zebra"}
	got := csgo.SortWords(w)

	th.AssertEqual(t, got, want)

}