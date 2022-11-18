package sort_test

import (
	csgo "github.com/cpustejovsky/customsortgo/sort"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSortForStrings(t *testing.T) {
	w := []string{"cat", "albatross", "dolphin", "bee", "zebra", "aardvark"}
	want := []string{"aardvark", "albatross", "bee", "cat", "dolphin", "zebra"}
	got := csgo.NewSortedStrings(w)
	assert.Equal(t, want, got)
}

func TestReverseSortForStrings(t *testing.T) {
	w := []string{"cat", "albatross", "dolphin", "bee", "zebra", "aardvark"}
	want := []string{"zebra", "dolphin", "cat", "bee", "albatross", "aardvark"}
	got := csgo.NewReverseSortedStrings(w)
	assert.Equal(t, want, got)
}

func TestSortInts(t *testing.T) {
	n := []int{1, 3, 5, 2, 7, 6, 4}
	want := []int{1, 2, 3, 4, 5, 6, 7}
	got := csgo.NewSortedIntegers(n)
	assert.Equal(t, want, got)
}

func TestReverseSortInts(t *testing.T) {
	n := []int{1, 3, 5, 2, 7, 6, 4}
	want := []int{7, 6, 5, 4, 3, 2, 1}
	got := csgo.NewReverseSortedIntegers(n)
	assert.Equal(t, want, got)
}

func TestSortFloat64s(t *testing.T) {
	n := []float64{1.1, 3.3, 5.5, 2.2, 7.7, 6.6, 4.4}
	want := []float64{1.1, 2.2, 3.3, 4.4, 5.5, 6.6, 7.7}
	got := csgo.NewSortedFloats(n)
	assert.Equal(t, want, got)
}

func TestReverseSortFloat64s(t *testing.T) {
	n := []float64{1.1, 3.3, 5.5, 2.2, 7.7, 6.6, 4.4}
	want := []float64{7.7, 6.6, 5.5, 4.4, 3.3, 2.2, 1.1}
	got := csgo.NewReverseSortedFloats(n)
	assert.Equal(t, want, got)
}
