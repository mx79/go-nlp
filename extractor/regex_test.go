package extractor

import "testing"

var valGetEntities2 = []struct {
	sentence string
	want     []string
}{
	{"A word is missing", []string{}},
	{"A word is missing", []string{}},
	{"A word is missing", []string{}},
}

func TestRegexExtractor_GetEntities(t *testing.T) {
	ext := NewRegexExtractor("")
	for _, test := range valGetEntities2 {
		got := ext.GetEntities(test.sentence)
		for idx := range got {
			if got[idx] != test.want[idx] {
				t.Fatalf("=> Got: %v\n=> Want: %v", got, test.want)
			}
		}
	}
}

var valGetSentences2 = []struct {
	strSlice []string
	want     []string
}{
	{[]string{}, []string{}},
	{[]string{}, []string{}},
	{[]string{}, []string{}},
}

func TestRegexExtractor_GetSentences(t *testing.T) {
	ext := NewRegexExtractor("")
	for _, test := range valGetSentences2 {
		got := ext.GetSentences(test.strSlice)
		for idx := range got {
			if got[idx] != test.want[idx] {
				t.Fatalf("=> Got: %v\n=> Want: %v", got, test.want)
			}
		}
	}
}
