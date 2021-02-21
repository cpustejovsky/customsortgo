package customsortgo

import (
	"fmt"
	"reflect"
	"sort"
)

//TurnToSortInterface takes a parameter
//if paramter is a slice of strings, ints, or float64s,
//it returns that slice converted to a sort.FooSlice type
//else it returns an error
func TurnToSortInterface(list interface{}) (sort.Interface, error) {
	switch reflect.TypeOf(list).String() {
	case "[]string":
		return sort.StringSlice(list.([]string)), nil
	case "[]int":
		return sort.IntSlice(list.([]int)), nil
	case "[]float64":
		return sort.Float64Slice(list.([]float64)), nil
	default:
		err := fmt.Errorf("Provided list was not made up of string, int, or float64 types. User provided list of type %v", reflect.TypeOf(list).String())
		return nil, err
	}
}

//TurnFromSortInterface takes a parameter
//if paramter is a slice of strings, ints, or float64s,
//it returns that slice converted to a sort.FooSlice type
//else it returns an error
func TurnFromSortInterface(list interface{}) (interface{}, error) {
	switch reflect.TypeOf(list).String() {
	case "sort.StringSlice":
		var s []string
		for _, v := range list.(sort.StringSlice) {
			s = append(s, v)
		}
		return s, nil
	case "sort.IntSlice":
		var s []int
		for _, v := range list.(sort.IntSlice) {
			s = append(s, v)
		}
		return s, nil
	case "sort.Float64Slice":
		var s []float64
		for _, v := range list.(sort.Float64Slice) {
			s = append(s, v)
		}
		return s, nil
	default:
		err := fmt.Errorf("Provided sort.Interface was not either types sort.StringSlice, sort.IntSlice, sort.Float64Slice. User provided list of type %v", reflect.TypeOf(list).String())
		return nil, err
	}
}

//Sort takes a slice and returns the items in a sorted order
func Sort(list interface{}) (interface{}, error) {
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

//ReverseSort takes a slice and returns the items in a reverse sorted order
func ReverseSort(list interface{}) (interface{}, error) {
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
