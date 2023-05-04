package clean_test

import (
	"github.com/mx79/go-nlp/clean"
	"testing"
)

func TestStopwords(t *testing.T) {
	examples := []struct {
		stopwords *clean.Stopwords
		input     string
		expected  string
	}{
		{
			stopwords: clean.NewStopwords(clean.EN),
			input:     "the quick brown fox jumped over the lazy dog.",
			expected:  "quick brown fox jumped lazy dog.",
		},
		{
			stopwords: clean.NewStopwords(clean.FR),
			input:     "le chat est sur le tapis.",
			expected:  "chat tapis.",
		},
		{
			stopwords: clean.NewStopwords(clean.ES),
			input:     "el perro corre por el parque.",
			expected:  "perro corre parque.",
		},
	}

	for _, ex := range examples {
		output := ex.stopwords.Stop(ex.input)
		if output != ex.expected {
			t.Errorf("Stop(%q) with result %q; expected %q", ex.input, output, ex.expected)
		}
	}
}
