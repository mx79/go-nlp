package extractor

import (
	"encoding/json"
	"github.com/mx79/go-nlp/utils"
	"io"
	"os"
	"regexp"
)

// RegexExtractor
type RegexExtractor map[string][]string

// NewRegexExtractor
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
