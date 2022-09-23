package extractor

import (
	_ "embed"
	"encoding/json"
	"github.com/mx79/go-nlp/utils"
	"io"
	"os"
	"regexp"
)

// Extractor
type Extractor interface {
	GetEntity()
	GetSentences()
}

//go:embed regex.json
var regexBytes []byte

// regex
var regex = loadRegex(regexBytes)

// DefaultRegexExtractor object
// In information retrieval, sometime we want to extract entities
// from a sentence or a text, that is where this object can be useful
// by implementing regular expression extraction from a map
type DefaultRegexExtractor map[string][]string

// loadRegex loads the map that contains the stopwords in many languages
func loadRegex(b []byte) DefaultRegexExtractor {
	var dict DefaultRegexExtractor
	utils.Check(json.Unmarshal(b, &dict))
	return dict
}

// NewDefaultRegexExtractor instantiates a new DefaultRegexExtractor object
func NewDefaultRegexExtractor() DefaultRegexExtractor {
	// Return an instance of RegexExtractor
	return regex
}

// GetEntity allows us to get back any entity and
// their corresponding matching pattern from our RegexExtractor dict
func (ext DefaultRegexExtractor) GetEntity(s string) map[string]interface{} {
	res := make(map[string]interface{})
	var patternList []string
	for k, v := range ext {
		for _, pattern := range v {
			re := regexp.MustCompile(pattern)
			if match := re.FindString(s); match != "" {
				patternList = append(patternList, match)
				patternList = utils.SortedSet(patternList)
				if len(patternList) > 1 {
					res[k] = patternList
				} else {
					res[k] = match
				}
			}
		}
		patternList = []string{}
	}
	return res
}

// GetSentences allows us to get back any sentences
// that match a pattern from our RegexExtractor dict
func (ext DefaultRegexExtractor) GetSentences(slice []string) []string {
	var res []string
	for _, v := range ext {
		for _, pattern := range v {
			for _, val := range slice {
				re := regexp.MustCompile(pattern)
				if match := re.FindString(val); match != "" {
					res = append(res, val)
				}
			}
		}
	}
	return res
}

// RegexExtractor object
// In information retrieval, sometime we want to extract entities
// from a sentence or a text, that is where this object can be useful
// by implementing regular expression extraction from a map
type RegexExtractor map[string][]string

// NewRegexExtractor instantiates a new RegexExtractor object
func NewRegexExtractor(regexFilePath string) RegexExtractor {
	// Opening the extractor.json file from the datapath
	jsonFile, err := os.Open(regexFilePath)
	utils.Check(err)
	defer jsonFile.Close()
	// Reading JSON file
	byteValue, _ := io.ReadAll(jsonFile)
	var dict map[string][]string
	utils.Check(json.Unmarshal(byteValue, &dict))
	// Return an instance of RegexExtractor
	return dict
}

// GetEntity allows us to get back any entity and
// their corresponding matching pattern from our RegexExtractor dict
func (ext RegexExtractor) GetEntity(s string) map[string]interface{} {
	res := make(map[string]interface{})
	var patternList []string
	for k, v := range ext {
		for _, pattern := range v {
			re := regexp.MustCompile(pattern)
			if match := re.FindString(s); match != "" {
				patternList = append(patternList, match)
				patternList = utils.SortedSet(patternList)
				if len(patternList) > 1 {
					res[k] = patternList
				} else {
					res[k] = match
				}
			}
		}
		patternList = []string{}
	}
	return res
}

// GetSentences allows us to get back any sentences
// that match a pattern from our RegexExtractor dict
func (ext RegexExtractor) GetSentences(slice []string) []string {
	var res []string
	for _, v := range ext {
		for _, pattern := range v {
			for _, val := range slice {
				re := regexp.MustCompile(pattern)
				if match := re.FindString(val); match != "" {
					res = append(res, val)
				}
			}
		}
	}
	return res
}
