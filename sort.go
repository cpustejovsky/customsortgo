package customsortgo

import (
	"fmt"
	"reflect"
	"sort"
)

type IncorrectTypeError struct {
	list any
}

func (e *IncorrectTypeError) Error() string {
	return fmt.Sprintf("provided sort.Interface was not either types sort.StringSlice, sort.IntSlice, sort.Float64Slice; user provided list of type %s", reflect.TypeOf(e.list).String())
}

// TurnToSortInterface takes a parameter
// if parameter is a slice of strings, ints, or float64s,
// it returns that slice converted to a sort.FooSlice type
// else it returns an error
func TurnToSortInterface(list any) (sort.Interface, error) {
	switch list.(type) {
	case []string:
		return sort.StringSlice(list.([]string)), nil
	case []int:
		return sort.IntSlice(list.([]int)), nil
	case []float64:
		return sort.Float64Slice(list.([]float64)), nil
	default:
		return nil, &IncorrectTypeError{list: list}
	}
}

// TurnFromSortInterface takes a parameter
// if parameter is type sort.StringSlice, sort.IntSlice, or sort.Float64Slice,
// it returns that slice converted to a []string,[]int, []float64 type, respectively
// else it returns an error
func TurnFromSortInterface(list any) (any, error) {
	switch list.(type) {
	case sort.StringSlice:
		var s []string
		for _, v := range list.(sort.StringSlice) {
			s = append(s, v)
		}
		return s, nil
	case sort.IntSlice:
		var s []int
		for _, v := range list.(sort.IntSlice) {
			s = append(s, v)
		}
		return s, nil
	case sort.Float64Slice:
		var s []float64
		for _, v := range list.(sort.Float64Slice) {
			s = append(s, v)
		}
		return s, nil
	default:
		return nil, &IncorrectTypeError{list: list}
	}
}

// Sort takes a slice and returns the items in a sorted order
func Sort(list any) (any, error) {
	s, err := TurnToSortInterface(list)
	if err != nil {
		return nil, err
	}
	sort.Sort(s)
	ns, err := TurnFromSortInterface(s)
	if err != nil {
		return nil, err
	}
	return ns, nil
}

// ReverseSort takes a slice and returns the items in a reverse sorted order
func ReverseSort(list any) (any, error) {
	s, err := TurnToSortInterface(list)
	if err != nil {
		return nil, err
	}
	sort.Sort(sort.Reverse(s))
	ns, err := TurnFromSortInterface(s)
	if err != nil {
		return nil, err
	}
	return ns, nil
}
