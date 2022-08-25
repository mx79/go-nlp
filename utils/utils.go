package utils

import (
	"os"
	"sort"
	"strings"
)

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
func SliceContains(slice []string, value string) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}

// SliceDeleteItem The function that
func SliceDeleteItem(slice []string, value string) []string {
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

// SortedWordSet The function that
func SortedWordSet(slice []string) []string {
	var newSlice []string
	for _, val := range slice {
		if !SliceContains(newSlice, val) {
			newSlice = append(newSlice, val)
		}
	}
	sort.Strings(newSlice)
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
