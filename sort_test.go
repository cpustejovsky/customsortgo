package customsortgo_test

import (
	csgo "github.com/cpustejovsky/customsortgo"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSortForStrings(t *testing.T) {
	w := []string{"cat", "albatross", "dolphin", "bee", "zebra", "aardvark"}
	want := []string{"aardvark", "albatross", "bee", "cat", "dolphin", "zebra"}
	got, err := csgo.Sort(w)
	assert.Nil(t, err)
	assert.Equal(t, want, got)
}

func TestReverseSortForStrings(t *testing.T) {
	w := []string{"cat", "albatross", "dolphin", "bee", "zebra", "aardvark"}
	want := []string{"zebra", "dolphin", "cat", "bee", "albatross", "aardvark"}
	got, err := csgo.ReverseSort(w)
	assert.Nil(t, err)
	assert.Equal(t, want, got)
}

func TestSortInts(t *testing.T) {
	n := []int{1, 3, 5, 2, 7, 6, 4}
	want := []int{1, 2, 3, 4, 5, 6, 7}
	got, err := csgo.Sort(n)
	assert.Nil(t, err)
	assert.Equal(t, want, got)
}

func TestReverseSortInts(t *testing.T) {
	n := []int{1, 3, 5, 2, 7, 6, 4}
	want := []int{7, 6, 5, 4, 3, 2, 1}
	got, err := csgo.ReverseSort(n)
	assert.Nil(t, err)
	assert.Equal(t, want, got)
}

func TestSortFloat64s(t *testing.T) {
	n := []float64{1.1, 3.3, 5.5, 2.2, 7.7, 6.6, 4.4}
	want := []float64{1.1, 2.2, 3.3, 4.4, 5.5, 6.6, 7.7}
	got, err := csgo.Sort(n)
	assert.Nil(t, err)
	assert.Equal(t, want, got)
}

func TestReverseSortFloat64s(t *testing.T) {
	n := []float64{1.1, 3.3, 5.5, 2.2, 7.7, 6.6, 4.4}
	want := []float64{7.7, 6.6, 5.5, 4.4, 3.3, 2.2, 1.1}
	got, err := csgo.ReverseSort(n)
	assert.Nil(t, err)
	assert.Equal(t, want, got)
}

type testStruct struct {
	Foo string
}

func TestSortError(t *testing.T) {
	ts := testStruct{Foo: "Bar"}
	c := []testStruct{ts, ts, ts}
	_, err := csgo.Sort(c)
	check := &csgo.IncorrectTypeError{}
	assert.ErrorAs(t, err, &check)
}
