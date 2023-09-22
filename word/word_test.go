package word_test

import (
	csgo "github.com/cpustejovsky/customsortgo/word"
	"testing"
)

func TestChars_Len(t *testing.T) {
	var testChars csgo.Chars = []rune("hello, world!")
	want := 13
	got := testChars.Len()
	if want != got {
		t.Errorf("got:\n%v\nwanted:\n%v\n", got, want)
	}
}

func TestChars_Swap(t *testing.T) {
	var testChars csgo.Chars = []rune("hello, world!")
	want := "ehllo, world!"
	testChars.Swap(0, 1)
	got := string(testChars)
	if want != got {
		t.Errorf("got:\n%v\nwanted:\n%v\n", got, want)
	}
}

func TestChars_Less(t *testing.T) {
	var testChars csgo.Chars = []rune("hello, world!")
	gotLessFirst := testChars.Less(0, 1)
	if gotLessFirst {
		t.Error("should return false")
	}
	testChars.Swap(0, 1)
	gotLessSecond := testChars.Less(0, 1)
	if !gotLessSecond {
		t.Error("should return true")
	}
}

func TestSortWord(t *testing.T) {
	w := "cbadegf"
	want := "abcdefg"
	got := csgo.SortWord(w)
	if want != got {
		t.Errorf("got:\n%v\nwanted:\n%v\n", got, want)
	}
}

func TestReverseSortWord(t *testing.T) {
	w := "cbadegf"
	want := "gfedcba"
	got := csgo.ReverseSortWord(w)
	if want != got {
		t.Errorf("got:\n%v\nwanted:\n%v\n", got, want)
	}
}
