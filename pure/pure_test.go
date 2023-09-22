package pure_test

import (
	"github.com/cpustejovsky/customsortgo/pure"
	"testing"
)

func TestNewSortForStrings(t *testing.T) {
	w := []string{"cat", "albatross", "dolphin", "bee", "zebra", "aardvark"}
	want := []string{"aardvark", "albatross", "bee", "cat", "dolphin", "zebra"}
	got := pure.NewSortedStrings(w)
	if !equalStringSlices(want, got) {
		t.Errorf("got:\n%v\nwanted:\n%v\n", got, want)
	}
	if equalStringSlices(w, got) {
		t.Errorf("orignal:\n%v\n and output:\n%v\nshould not be equal\n", w, got)
	}
}

func TestNewReverseSortForStrings(t *testing.T) {
	w := []string{"cat", "albatross", "dolphin", "bee", "zebra", "aardvark"}
	want := []string{"zebra", "dolphin", "cat", "bee", "albatross", "aardvark"}
	got := pure.NewReverseSortedStrings(w)
	if !equalStringSlices(want, got) {
		t.Errorf("got:\n%v\nwanted:\n%v\n", got, want)
	}
	if equalStringSlices(w, got) {
		t.Errorf("orignal:\n%v\n and output:\n%v\nshould not be equal\n", w, got)
	}
}

func TestNewSortInts(t *testing.T) {
	n := []int{1, 3, 5, 2, 7, 6, 4}
	want := []int{1, 2, 3, 4, 5, 6, 7}
	got := pure.NewSortedIntegers(n)
	if !equalIntSlices(want, got) {
		t.Errorf("got:\n%v\nwanted:\n%v\n", got, want)
	}
	if equalIntSlices(n, got) {
		t.Errorf("orignal:\n%v\n and output:\n%v\nshould not be equal\n", n, got)
	}
}

func TestNewReverseSortInts(t *testing.T) {
	n := []int{1, 3, 5, 2, 7, 6, 4}
	want := []int{7, 6, 5, 4, 3, 2, 1}
	got := pure.NewReverseSortedIntegers(n)
	if !equalIntSlices(want, got) {
		t.Errorf("got:\n%v\nwanted:\n%v\n", got, want)
	}
	if equalIntSlices(n, got) {
		t.Errorf("orignal:\n%v\n and output:\n%v\nshould not be equal\n", n, got)
	}
}

func TestNewSortFloat64s(t *testing.T) {
	n := []float64{1.1, 3.3, 5.5, 2.2, 7.7, 6.6, 4.4}
	want := []float64{1.1, 2.2, 3.3, 4.4, 5.5, 6.6, 7.7}
	got := pure.NewSortedFloats(n)
	if !equalFloatSlices(want, got) {
		t.Errorf("got:\n%v\nwanted:\n%v\n", got, want)
	}
	if equalFloatSlices(n, got) {
		t.Errorf("orignal:\n%v\n and output:\n%v\nshould not be equal\n", n, got)
	}
}

func TestNewReverseSortFloat64s(t *testing.T) {
	n := []float64{1.1, 3.3, 5.5, 2.2, 7.7, 6.6, 4.4}
	want := []float64{7.7, 6.6, 5.5, 4.4, 3.3, 2.2, 1.1}
	got := pure.NewReverseSortedFloats(n)
	if !equalFloatSlices(want, got) {
		t.Errorf("got:\n%v\nwanted:\n%v\n", got, want)
	}
	if equalFloatSlices(n, got) {
		t.Errorf("orignal:\n%v\n and output:\n%v\nshould not be equal\n", n, got)
	}
}

func equalStringSlices(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, s := range a {
		if s != b[i] {
			return false
		}
	}
	return true
}

func equalIntSlices(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, s := range a {
		if s != b[i] {
			return false
		}
	}
	return true
}

func equalFloatSlices(a, b []float64) bool {
	if len(a) != len(b) {
		return false
	}
	for i, s := range a {
		if s != b[i] {
			return false
		}
	}
	return true
}
