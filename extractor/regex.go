package extractor

import (
	"log"
	"regexp"
)

// RegexExtractor object:
//
// In information retrieval, sometime we want to extract entities
// from a sentence or a text, that is where this object can be useful
// by implementing regular expression extraction from a pattern.
type RegexExtractor struct {
	Pattern *regexp.Regexp
	Flags   map[RegexFlag]bool
}

// NewRegexExtractor instantiates a new RegexExtractor object.
//
// In information retrieval, sometime we want to extract entities
// from a sentence or a text, that is where this object can be useful
// by implementing regular expression extraction from a pattern.
func NewRegexExtractor(pattern string, flags ...RegexFlag) *RegexExtractor {
	for _, f := range flags {
		if _, b := flagMap[f]; b {
			flagMap[f] = true
		} else {
			log.Fatalf("Unidentified flag named: %v\n"+
				"You should choose one from the go-nlp lib constants, ex: IGNORECASE", f)
		}
	}

	re := adjustPattern(pattern, flagMap)

	return &RegexExtractor{
		Pattern: re,
		Flags:   flagMap,
	}
}

// GetEntity extracts any match with the fixed pattern and flags.
//
// It returns a slice of match.
func (ext *RegexExtractor) GetEntity(s string) (res []string) {
	for _, match := range ext.Pattern.FindAllString(s, -1) {
		if match != "" {
			res = append(res, match)
		}
	}

	return
}

// GetSentences allows us to get back any sentences that contains a match with our pattern.
func (ext *RegexExtractor) GetSentences(slice []string) (res []string) {
	for _, val := range slice {
		if match := ext.Pattern.FindString(val); match != "" {
			res = append(res, val)
		}
	}

	return
}
