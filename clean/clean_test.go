package clean_test

import (
	"testing"

	"github.com/mx79/go-nlp/clean"
)

func TestRemovePunctuation(t *testing.T) {
	examples := []struct {
		input    string
		expected string
	}{
		{
			input:    "Hello, World!",
			expected: "Hello World",
		},
		{
			input:    "The quick brown fox jumped over the lazy dog.",
			expected: "The quick brown fox jumped over the lazy dog",
		},
		{
			input:    "Mary had a little lamb, its fleece was white as snow.",
			expected: "Mary had a little lamb its fleece was white as snow",
		},
		{
			input:    "This is a test.",
			expected: "This is a test",
		},
	}

	for _, ex := range examples {
		output := clean.RemovePunctuation(ex.input)
		if output != ex.expected {
			t.Errorf("RemovePunctuation(%q) = %q; expected %q", ex.input, output, ex.expected)
		}
	}
}

func TestRemoveAccent(t *testing.T) {
	examples := []struct {
		input    string
		expected string
	}{
		{
			input:    "héllo wórld",
			expected: "hello world",
		},
		{
			input:    "à bientôt",
			expected: "a bientot",
		},
		{
			input:    "équipe de France",
			expected: "equipe de France",
		},
		{
			input:    "Beyoncé Knowles",
			expected: "Beyonce Knowles",
		},
		{
			input:    "Crème brûlée",
			expected: "Creme brulee",
		},
	}

	for _, ex := range examples {
		output := clean.RemoveAccent(ex.input)
		if output != ex.expected {
			t.Errorf("RemoveAccent(%q) = %q; expected %q", ex.input, output, ex.expected)
		}
	}
}

func TestTokenize(t *testing.T) {
	examples := []struct {
		input    string
		expected []string
		withPunc bool
	}{
		{
			input:    "Hi! How are you?",
			expected: []string{"Hi", "!", "How", "are", "you", "?"},
			withPunc: true,
		},
		{
			input:    "The quick brown fox jumped over the lazy dog.",
			expected: []string{"The", "quick", "brown", "fox", "jumped", "over", "the", "lazy", "dog", "."},
			withPunc: true,
		},
		{
			input:    "Mary had a little lamb, its fleece was white as snow.",
			expected: []string{"Mary", "had", "a", "little", "lamb", ",", "its", "fleece", "was", "white", "as", "snow", "."},
			withPunc: true,
		},
		{
			input:    "  This    is a     test.  ",
			expected: []string{"This", "is", "a", "test", "."},
			withPunc: true,
		},
	}

	for _, ex := range examples {
		output := clean.Tokenize(ex.input, ex.withPunc)
		if !equalSlices(output, ex.expected) {
			t.Errorf("Tokenize(%q, %t) = %v; expected %v", ex.input, ex.withPunc, output, ex.expected)
		}
	}
}

// equalSlices returns true if two string slices are equal.
func equalSlices(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
