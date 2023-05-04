package distance_test

import (
	"github.com/mx79/go-nlp/distance"
	"testing"
)

func TestWordErrorRate(t *testing.T) {
	examples := []struct {
		a        string
		b        string
		expected float64
	}{
		{
			a:        "The quick brown fox jumps over the lazy dog.",
			b:        "The quick brown fox jumps over the lazy dog.",
			expected: 0,
		},
		{
			a:        "The quick brown fox jumps over the lazy dog.",
			b:        "The quick brown fox jumps over the lazy cat.",
			expected: 0.1,
		},
		{
			a:        "The rain in Spain stays mainly on the plain.",
			b:        "In Spain, it rains mainly on the plain.",
			expected: 0.5,
		},
	}

	for _, ex := range examples {
		output := distance.WordErrorRate(ex.a, ex.b)
		if output != ex.expected {
			t.Errorf("WordErrorRate(%q, %q) = %f; expected %f", ex.a, ex.b, output, ex.expected)
		}
	}
}
