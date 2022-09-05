package clean

import (
	"github.com/mx79/go-nlp/utils"
	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
	"strings"
	"unicode"
)

// RemovePunctuation removes punctuation in a sentence
// The punctuation set does not include the "-" char
// and the simple quote as their used into some languages
func RemovePunctuation(s string) string {
	punctuation := "!@#$%^&*()[]{}<>_+?:.,;"
	for _, c := range punctuation {
		s = strings.Replace(s, string(c), "", -1)
	}
	return s
}

// RemoveAccent The function that allows you to remove the accents in a sentence
func RemoveAccent(s string) string {
	t := transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC)
	output, _, err := transform.String(t, s)
	utils.Check(err)
	return output
}

// Lower The function that put words of a sentence in lowercase mode
func Lower(s string) string {
	return strings.ToLower(s)
}

// Tokenize
func Tokenize(s string) []string {
	return strings.Split(strings.Trim(s, " "), " ")
}
