package extractor

import (
	"encoding/json"
	"github.com/mx79/go-nlp/clean"
	"github.com/mx79/go-nlp/utils"
	"io"
	"os"
	"regexp"
)

// RegexExtractor
type RegexExtractor struct {
	RegexDict map[string][]string
}

// NewRegexExtractor
func NewRegexExtractor(regexFilePath string) *RegexExtractor {
	// Opening the extractor.json file from the datapath
	jsonFile, err := os.Open(regexFilePath)
	if err != nil {
		panic(err)
	}
	defer jsonFile.Close()
	// Reading JSON file
	byteValue, _ := io.ReadAll(jsonFile)
	var dict map[string][]string
	if err = json.Unmarshal(byteValue, &dict); err != nil {
		panic(err)
	}
	// Return an instance of RegexExtractor
	return &RegexExtractor{
		RegexDict: dict,
	}
}

// GetEntity
func (ext *RegexExtractor) GetEntity(s string) map[string]interface{} {
	cleanedText := clean.RemovePunctuation(s)
	res := make(map[string]interface{})
	var patternList []string
	for k, v := range ext.RegexDict {
		for _, pattern := range v {
			re := regexp.MustCompile(pattern)
			if match := re.FindString(cleanedText); match != "" {
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
