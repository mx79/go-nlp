package extractor

import (
	"encoding/json"
	"github.com/mx79/go-nlp/clean"
	"github.com/mx79/go-nlp/utils"
	"io"
	"os"
	"regexp"
)

// EntityExtractor
type EntityExtractor struct {
	RegexDict map[string][]string
}

// NewEntityExtractor
func NewEntityExtractor(regexFilePath string) *EntityExtractor {
	// Opening the extractor.json file from the datapath
	jsonFile, e2 := os.Open(regexFilePath)
	if e2 != nil {
		panic(e2)
	}
	defer jsonFile.Close()
	// Reading JSON file
	byteValue, _ := io.ReadAll(jsonFile)
	var dict map[string][]string
	e3 := json.Unmarshal(byteValue, &dict)
	if e3 != nil {
		panic(e3)
	}
	// Return an instance of EntityExtractor
	return &EntityExtractor{
		RegexDict: dict,
	}
}

// GetEntity
func (ext *EntityExtractor) GetEntity(s string) map[string]interface{} {
	cleanedText := clean.RemovePunctuation(s)
	res := make(map[string]interface{})
	var patternList []string
	for k, v := range ext.RegexDict {
		for _, pattern := range v {
			re := regexp.MustCompile(pattern)
			if match := re.FindString(cleanedText); match != "" {
				patternList = append(patternList, match)
				patternList = utils.SortedWordSet(patternList)
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
