package utils

import (
	"sort"
	"strings"
)

type (
	// Global is a generic type instantiating all Golang types
	Global interface {
		string | int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64 | bool
	}

	// SubGlobal is a generic type instantiating the 3 basic types of Golang
	SubGlobal interface {
		string | int | float64
	}
)

// ListToStr transforms a slice of string into a string
func ListToStr(strSlice []string) string {
	return strings.Join(strSlice, " ")
}

// SliceContains checks if a given slice contains an element
// return true if the elem is in the slice else false
func SliceContains[T Global](slice []T, value T) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}

// SliceDeleteItem deletes the selected element in a slice
// It returns the slice without the element, but keep in mind
// that for the moment this function must be used with a set of values
// because it is only deleting the first match with the input element
// If there are many same elements in the slice, it will only delete one.
func SliceDeleteItem[T Global](slice []T, value T) []T {
	idxToDel := 0
	for _, item := range slice {
		if item == value {
			break
		}
		idxToDel++
	}
	if idxToDel >= len(slice) {
		return slice
	} else {
		newSlice := append(slice[:idxToDel], slice[(idxToDel+1):]...)
		return newSlice
	}
}

// MapContains checks if a map contains an element
// return true if the elem in the map else false
func MapContains[T Global](m map[T]T, value T) bool {
	if _, ok := m[value]; ok {
		return true
	}
	return false
}

// Set is a function that mimics the behavior of a set by removing
// duplicate values in a slice, because the set type does not exist in Golang
func Set[T Global](slice []T) []T {
	var newSlice []T
	for _, val := range slice {
		if !SliceContains(newSlice, val) {
			newSlice = append(newSlice, val)
		}
	}
	return newSlice
}

// Sorted sorts the values of a slice of type:
// string, int or float64
func Sorted[T SubGlobal](slice []T) {
	switch s := any(slice).(type) {
	case []string:
		sort.Strings(s)
	case []int:
		sort.Ints(s)
	case []float64:
		sort.Float64s(s)
	}
}

// SortedSet uses both Set and Sorted
// functions as explained in its name
// But it only works for string, int and float64 types
func SortedSet[T SubGlobal](slice []T) []T {
	newSlice := Set(slice)
	Sorted(newSlice)
	return newSlice
}

// Check is a simple function to remove the redundant
// use of the Golang error detection system
func Check(e error) {
	if e != nil {
		panic(e)
	}
}
