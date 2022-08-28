package utils

import (
	"os"
	"sort"
	"strings"
)

type Global interface {
	string | int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64
}

type SubGlobal interface {
	string | int | float64
}

// ListToStr The function that transforms a list into a string
func ListToStr(strSlice []string) string {
	return strings.Join(strSlice, " ")
}

// SplitOnConj The function that allows to separate a sentence on given conjunctions
func SplitOnConj(s string) []string {
	conjList := []string{" mais ", " donc ", " or ", " car ", " parce que ", " en fait "}
	for _, conj := range conjList {
		if strings.Contains(s, conj) {
			s = strings.Replace(s, conj, "#", -1)
		}
	}
	return strings.Split(s, "#")
}

// SliceContains The function that
func SliceContains[T Global](slice []T, value T) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}

// SliceDeleteItem The function that
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

// Set The function that
func Set[T Global](slice []T) []T {
	var newSlice []T
	for _, val := range slice {
		if !SliceContains(newSlice, val) {
			newSlice = append(newSlice, val)
		}
	}
	return newSlice
}

// Sorted The function that
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

// SortedSet The function that
func SortedSet[T SubGlobal](slice []T) []T {
	newSlice := Set(slice)
	Sorted(newSlice)
	return newSlice
}

// Check
func Check(e error) {
	if e != nil {
		panic(e)
	}
}

// Write
func Write(text string, file *os.File) {
	if _, err := file.WriteString(text); err != nil {
		panic(err)
	}
}

// Read
func Read(filename string) string {
	data, err := os.ReadFile(filename)
	Check(err)
	return string(data)
}
