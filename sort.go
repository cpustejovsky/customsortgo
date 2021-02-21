package customsortgo

import (
	"errors"
	"reflect"
	"sort"
)

//DetermineSortedType takes a parameter
//if paramter is a slice of strings, ints, or float64s,
//it returns that slice converted to a sort.FooSlice type
//else it returns an error
func DetermineSortedType(list interface{}) (sort.Interface, error) {
	switch reflect.TypeOf(list).String() {
	case "[]string":
		return sort.StringSlice(list.([]string)), nil
	case "[]int":
		return sort.IntSlice(list.([]int)), nil
	case "[]float64":
		return sort.Float64Slice(list.([]float64)), nil
	default:
		err := errors.New("Provided list was not made up of string, int, or float64 types")
		return nil, err
	}
}

//Sort takes a slice and returns the items in a sorted order
func Sort(list interface{}) (sort.Interface, error) {
	s, err := DetermineSortedType(list)
	if err != nil {
		return nil, err
	}
	sort.Sort(s)
	return s, nil
}

//ReverseSort takes a slice and returns the items in a reverse sorted order
func ReverseSort(list interface{}) (sort.Interface, error) {
	s, err := DetermineSortedType(list)
	if err != nil {
		return nil, err
	}
	sort.Sort(sort.Reverse(s))
	return s, nil
}
