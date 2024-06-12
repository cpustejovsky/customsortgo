package count_test

import (
	"github.com/cpustejovsky/customsortgo/count"
	"testing"
)

func TestNewSorted(t *testing.T) {
	input := []int{5, 2, 4, 6, 1, 3}
	want := []int{1, 2, 3, 4, 5, 6}
	m := 6
	got := count.NewSorted(input, m)
	if !sameSlices(got, want) {
		t.Fatalf("got:\t%v\nwant:\t%v\n", got, want)
	}
}

func sameSlices[C comparable](a, b []C) bool {
	for i, element := range a {
		if element != b[i] {
			return false
		}
	}
	return true
}
