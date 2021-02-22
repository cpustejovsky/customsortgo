package customsortgo_test

import (
	"testing"

	csgo "github.com/cpustejovsky/customsortgo"
	th "github.com/cpustejovsky/customsortgo/testhelp"
)

func TestReverseString(t *testing.T) {
	w := "abdc"
	want := "cdba"
	got := csgo.ReverseString(w)
	th.AssertEqual(t, got, want)
}
