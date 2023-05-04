package clean_test

import (
	"github.com/mx79/go-nlp/clean"
	"testing"
)

func TestPurge(t *testing.T) {
	examples := []struct {
		purger   *clean.TextPurger
		input    string
		expected string
	}{
		{
			purger:   clean.NewTextPurger(clean.FR, true, true, true, true, true),
			input:    "Bonjour, comment allez-vous aujourd'hui ?",
			expected: "bonjour allez-vous aujourd'hui",
		},
		{
			purger:   clean.NewTextPurger(clean.EN, false, true, false, false, true),
			input:    "The quick brown fox jumped over the dog.",
			expected: "the quick brown fox jump over the dog.",
		},
		{
			purger:   clean.NewTextPurger(clean.EN, true, false, true, true, false),
			input:    "sells seashells by the seashore.",
			expected: "sells seashells seashore",
		},
	}

	for _, ex := range examples {
		output := ex.purger.Purge(ex.input)
		if output != ex.expected {
			t.Errorf("Purge(%q) with result %q; expected %q", ex.input, output, ex.expected)
		}
	}
}
