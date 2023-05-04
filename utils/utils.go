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

// ListToStr is a shortcut for the usage of `join` from strings methods.
//
// It is used to return a slice of string concatenated
// as one string without any ambiguity possible.
func ListToStr(strSlice []string) string {
	return strings.Join(strSlice, " ")
}

// SliceContains checks if a given slice contains an element.
//
// It returns true if the elem is in the slice else false.
func SliceContains[T Global](slice []T, value T) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}

	return false
}

// sliceDeleteOneItem deletes the selected element in a slice.
//
// It returns the slice without the indicated element, removed only one time.
func sliceDeleteOneItem[T Global](slice []T, value T) []T {
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

// SliceDeleteItem deletes the selected element recursively.
//
// It uses SliceDeleteOneItem until the item is no longer
// present in our slice
func SliceDeleteItem[T Global](slice []T, value T) []T {
	var count int
	for _, t := range slice {
		if t == value {
			count++
		}
	}
	if count != 0 {
		for {
			slice = sliceDeleteOneItem(slice, value)
			count--
			if count == 0 {
				break
			}
		}
	}

	return slice
}

// MapContains checks if a map contains an element.
//
// It returns true if the elem in the map else false
func MapContains[T Global](m map[T]T, value T) bool {
	_, ok := m[value]

	return ok
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
//
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
// functions as explained in its name.
//
// But it only works for string, int and float64 types.
func SortedSet[T SubGlobal](slice []T) []T {
	newSlice := Set(slice)
	Sorted(newSlice)

	return newSlice
}
