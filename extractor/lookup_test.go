package extractor

import (
	"testing"
)

func AssertPass(t *testing.T, got, want interface{}) {
	if got != want {
		t.Fatalf("=> Got: %v\n=> Want: %v", got, want)
	}
}

var valGetEntities = []struct {
	sentence string
	want     []string
}{
	{"A word is missing", []string{}},
	{"A word is missing", []string{}},
	{"A word is missing", []string{}},
}

func TestLookupExtractor_GetEntities(t *testing.T) {
	ext := NewLookupExtractor("regex.json")
	for _, test := range valGetEntities {
		got := ext.GetEntities(test.sentence)
		AssertPass(t, got, test.want)
	}
}

var valGetSentences = []struct {
	strSlice []string
	want     []string
}{
	{[]string{}, []string{}},
	{[]string{}, []string{}},
	{[]string{}, []string{}},
}

func TestLookupExtractor_GetSentences(t *testing.T) {
	ext := NewLookupExtractor("regex.json")
	for _, test := range valGetSentences {
		got := ext.GetSentences(test.strSlice)
		for idx := range got {
			if got[idx] != test.want[idx] {
				t.Fatalf("=> Got: %v\n=> Want: %v", got, test.want)
			}
		}
	}
}
