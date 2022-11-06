package extractor

import (
	"log"
	"regexp"
)

type RegexExtractor struct {
	EntityName string
	Pattern    *regexp.Regexp
	Flags      map[RegexFlag]bool
}

// NewRegexExtractor instantiates a new RegexExtractor object
// In information retrieval, sometime we want to extract entities
// from a sentence or a text, that is where this object can be useful
// by implementing regular expression extraction from a pattern
func NewRegexExtractor(entityName string, pattern string, flags ...RegexFlag) *RegexExtractor {
	// Implementing Flags
	var flagMap = map[RegexFlag]bool{
		IGNORECASE:   false,
		MULTILINE:    false,
		MATCHNEWLINE: false,
		UNGREEDY:     false,
	}
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
		EntityName: entityName,
		Pattern:    re,
		Flags:      flagMap,
	}
}

// GetEntity extracts any match with the fixed pattern and flags
// It returns a map with one entry , the entity name and the slice of match
func (ext *RegexExtractor) GetEntity(s string) map[string][]string {
	res := make(map[string][]string)
	for _, match := range ext.Pattern.FindAllString(s, -1) {
		if match != "" {
			res[ext.EntityName] = append(res[ext.EntityName], match)
		}
	}
	return res
}

// GetSentences allows us to get back any sentences that contains a match with our pattern
func (ext *RegexExtractor) GetSentences(slice []string) (res []string) {
	for _, val := range slice {
		if match := ext.Pattern.FindString(val); match != "" {
			res = append(res, val)
		}
	}
	return
}
