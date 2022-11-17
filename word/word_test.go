package word_test

import (
	csgo "github.com/cpustejovsky/customsortgo/word"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestChars_Len(t *testing.T) {
	var testChars csgo.Chars = []rune("hello, world!")
	wantLen := 13
	gotLen := testChars.Len()
	assert.Equal(t, wantLen, gotLen)
}

func TestChars_Swap(t *testing.T) {
	var testChars csgo.Chars = []rune("hello, world!")
	wantSwap := "ehllo, world!"
	testChars.Swap(0, 1)
	gotSwap := string(testChars)
	assert.Equal(t, wantSwap, gotSwap)
}

func TestChars_Less(t *testing.T) {
	var testChars csgo.Chars = []rune("hello, world!")
	gotLessFirst := testChars.Less(0, 1)
	assert.Equal(t, false, gotLessFirst)
	testChars.Swap(0, 1)
	gotLessSecond := testChars.Less(0, 1)
	assert.Equal(t, true, gotLessSecond)
}

func TestSortWord(t *testing.T) {
	w := "cbadegf"
	want := "abcdefg"
	got := csgo.SortWord(w)
	assert.Equal(t, want, got)
}

func TestReverseSortWord(t *testing.T) {
	w := "cbadegf"
	want := "gfedcba"
	got := csgo.ReverseSortWord(w)
	assert.Equal(t, want, got)
}
