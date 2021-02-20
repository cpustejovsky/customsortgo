package testhelp

import (
	"reflect"
	"testing"
)

func AssertEqual(t *testing.T, got, want interface{}) {
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got:\n%v\nwant\n%v\n", got, want)
	}
}
