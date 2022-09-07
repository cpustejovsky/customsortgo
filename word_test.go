package customsortgo_test

import (
	"github.com/stretchr/testify/assert"
	"testing"

	csgo "github.com/cpustejovsky/customsortgo"
)

var testChars csgo.Chars = []rune("hello, world!")

func TestChars(t *testing.T) {
	wantLen := 13
	gotLen := testChars.Len()
	assert.Equal(t, wantLen, gotLen)

	gotLessFirst := testChars.Less(0, 1)
	assert.Equal(t, false, gotLessFirst)

	wantSwap := "ehllo, world!"
	testChars.Swap(0, 1)
	gotSwap := string(testChars)
	assert.Equal(t, wantSwap, gotSwap)

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
