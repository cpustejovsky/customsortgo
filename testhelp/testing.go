package testhelp

import (
	"reflect"
	"strings"
	"testing"
)

// AssertEqual takes two interfaces and determines if they are equal
func AssertEqual(t *testing.T, got, want any) {
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got:\n%v\nwant\n%v\n", got, want)
	}
}

// ErrorContains checks if the error message in out contains the text in
// want.
//
// This is safe when out is nil. Use an empty string for want if you want to
// test that err is nil.
// pulled from: https://github.com/Teamwork/test/blob/7949982/test.go#L12-L25
func ErrorContains(out error, want string) bool {
	if out == nil {
		return want == ""
	}
	if want == "" {
		return false
	}
	return strings.Contains(out.Error(), want)
}
