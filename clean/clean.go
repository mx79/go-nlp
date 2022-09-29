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

// RemoveAccent allows you to remove the accents in a sentence
func RemoveAccent(s string) string {
	t := transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC)
	output, _, err := transform.String(t, s)
	utils.Check(err)
	return output
}

// Lower puts words of a sentence in lowercase mode
func Lower(s string) string {
	return strings.ToLower(s)
}

// Tokenize returns a sentence broken into tokens. Tokens are individual
// words as well as punctuation. For example, "Hi! How are you?" becomes
// []string{"Hi", "!", "How", "are", "you", "?"}.
func Tokenize(sent string, withPunct bool) []string {
	var tokens []string
	for _, w := range strings.Fields(sent) {
		var found []int
		for i, r := range w {
			if withPunct {
				switch r {
				case '\'', '"', ':', ';', '!', '?':
					found = append(found, i)
				// Handle case of currencies and fractional percents.
				case '.', ',':
					if i+1 < len(w) {
						switch w[i+1] {
						case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
							continue
						}
					}
					found = append(found, i)
					i++
				}
			}
		}
		if len(found) == 0 {
			tokens = append(tokens, w)
			continue
		}
		for i, j := range found {
			// If the token marker is not the first character in the
			// sentence, then include all characters leading up to
			// the prior found token.
			if j > 0 {
				if i == 0 {
					tokens = append(tokens, w[:j])
				} else if i-1 < len(found) {
					// Handle case where multiple tokens are
					// found in the same word.
					tokens = append(tokens, w[found[i-1]+1:j])
				}
			}
			// Append the token marker itself
			tokens = append(tokens, string(w[j]))
			// If we're on the last token marker, append all
			// remaining parts of the word.
			if i+1 == len(found) {
				tokens = append(tokens, w[j+1:])
			}
		}
	}
	return utils.SliceDeleteItem(tokens, "")
}
