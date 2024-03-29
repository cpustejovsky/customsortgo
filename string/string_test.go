package string_test

import (
	csgo "github.com/cpustejovsky/customsortgo/string"
	"testing"
)

func TestReverseString(t *testing.T) {
	w := "abdc"
	want := "cdba"
	got := csgo.ReverseString(w)
	if want != got {
		t.Errorf("got:\n%v\nwanted:\n%v\n", got, want)
	}
}

func TestReverseStringSlow(t *testing.T) {
	w := "abdc"
	want := "cdba"
	got := csgo.ReverseStringSlow(w)
	if want != got {
		t.Errorf("got:\n%v\nwanted:\n%v\n", got, want)
	}
}

var longStr = "A9Fn3WqFTqhyH8t&K$rPZ5NCsYCQgzPGT&7X%Acj#NB69YnyV!YLrNNCQq3s%fxs9okr5&XcT*GXMVEajc3RWo^XHx4Zx2A*5WCHjecbcLWCY2VAf6g6P7aeg^FZfBFq"

func BenchmarkReverseString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		csgo.ReverseString(longStr)
	}
}
func BenchmarkReverseStringSlow(b *testing.B) {
	for i := 0; i < b.N; i++ {
		csgo.ReverseStringSlow(longStr)
	}
}
