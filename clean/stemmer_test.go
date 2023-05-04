package clean_test

import (
	"github.com/mx79/go-nlp/clean"
	"testing"
)

func TestStem(t *testing.T) {
	examples := []struct {
		stemmer  *clean.Stemmer
		input    string
		expected string
	}{
		{
			stemmer:  clean.NewStemmer(clean.FR),
			input:    "autres",
			expected: "autr",
		},
		{
			stemmer:  clean.NewStemmer(clean.EN),
			input:    "jumping",
			expected: "jump",
		},
	}

	for _, ex := range examples {
		output := ex.stemmer.Stem(ex.input)
		if output != ex.expected {
			t.Errorf("Stem(%q) with result %q; expected %q", ex.input, output, ex.expected)
		}
	}
}
