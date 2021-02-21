package customsortgo_test

import (
	"sort"
	"testing"

	csgo "github.com/cpustejovsky/customsortgo"
	th "github.com/cpustejovsky/customsortgo/testhelp"
)

func TestSortForStrings(t *testing.T) {
	w := []string{"cat", "albatross", "dolphin", "bee", "zebra", "aardvark"}
	want := []string{"aardvark", "albatross", "bee", "cat", "dolphin", "zebra"}
	got, err := csgo.Sort(w)
	if err != nil {
		t.Error(err)
	} else {
		th.AssertEqual(t, got, sort.StringSlice(want))
	}
}

func TestReverseSortForStrings(t *testing.T) {
	w := []string{"cat", "albatross", "dolphin", "bee", "zebra", "aardvark"}
	want := []string{"zebra", "dolphin", "cat", "bee", "albatross", "aardvark"}
	got, err := csgo.ReverseSort(w)
	if err != nil {
		t.Error(err)
	} else {
		th.AssertEqual(t, got, sort.StringSlice(want))
	}
}

func TestSortInts(t *testing.T) {
	n := []int{1, 3, 5, 2, 7, 6, 4}
	want := []int{1, 2, 3, 4, 5, 6, 7}
	got, err := csgo.Sort(n)
	if err != nil {
		t.Error(err)
	} else {
		th.AssertEqual(t, got, sort.IntSlice(want))
	}
}

func TestReverseSortInts(t *testing.T) {
	n := []int{1, 3, 5, 2, 7, 6, 4}
	want := []int{7, 6, 5, 4, 3, 2, 1}
	got, err := csgo.ReverseSort(n)
	if err != nil {
		t.Error(err)
	} else {
		th.AssertEqual(t, got, sort.IntSlice(want))
	}
}

func TestSortFloat64s(t *testing.T) {
	n := []float64{1.1, 3.3, 5.5, 2.2, 7.7, 6.6, 4.4}
	want := []float64{1.1, 2.2, 3.3, 4.4, 5.5, 6.6, 7.7}
	got, err := csgo.Sort(n)
	if err != nil {
		t.Error(err)
	} else {
		th.AssertEqual(t, got, sort.Float64Slice(want))
	}
}

func TestReverseSortFloat64s(t *testing.T) {
	n := []float64{1.1, 3.3, 5.5, 2.2, 7.7, 6.6, 4.4}
	want := []float64{7.7, 6.6, 5.5, 4.4, 3.3, 2.2, 1.1}
	got, err := csgo.ReverseSort(n)
	if err != nil {
		t.Error(err)
	} else {
		th.AssertEqual(t, got, sort.Float64Slice(want))
	}
}

func TestSortError(t *testing.T) {
	c := "test"
	want := "Provided list was not made up of string, int, or float64 types. User provided list of type string"
	_, err := csgo.Sort(c)
	if !(th.ErrorContains(err, want)) {
		t.Errorf("got:\n%v\nwant\n%v\n", err, want)
	}
}
