package extractor

import (
	_ "embed"
	"encoding/json"
	"github.com/mx79/go-nlp/utils"
	"io"
	"log"
	"os"
	"regexp"
)

// Flags const, can be imported in the code and used
// to declare faster a LookupExtractor or RegexExtractor
const (
	IGNORECASE   = "IGNORECASE"
	MULTILINE    = "MULTILINE"
	MATCHNEWLINE = "MATCHNEWLINE"
	UNGREEDY     = "UNGREEDY"
)

// Extractor interface implements GetEntity and GetSentences methods
type Extractor interface {
	GetEntity()
	GetSentences()
}

// RegexFlag is an alias to string type,
// makes just more clear about what kind of string it is
type RegexFlag string

// LookupExtractor object
// In information retrieval, sometime we want to extract entities
// from a sentence or a text, that is where this object can be useful
// by implementing lookup expression extraction from a map
type LookupExtractor struct {
	LookupTable map[string][]string
	Flags       map[RegexFlag]bool
}

// NewLookupExtractor instantiates a new LookupExtractor object
func NewLookupExtractor(regexFilePath string, flags ...RegexFlag) *LookupExtractor {
	// Opening the provided json
	jsonFile, err := os.Open(regexFilePath)
	utils.Check(err)
	defer jsonFile.Close()
	// Reading JSON file
	byteValue, _ := io.ReadAll(jsonFile)
	var dict map[string][]string
	utils.Check(json.Unmarshal(byteValue, &dict))
	// Implementing Flags
	var flagMap = map[RegexFlag]bool{
		IGNORECASE:   false,
		MULTILINE:    false,
		MATCHNEWLINE: false,
		UNGREEDY:     false,
	}
	for _, f := range flags {
		if b := flagMap[f]; b {
			flagMap[f] = true
		} else {
			log.Fatalf("Unidentified flag named: %v\n"+
				"You should choose one from the go-nlp lib constants, ex: IGNORECASE", f)
		}
	}

	return &LookupExtractor{
		LookupTable: dict,
		Flags:       flagMap,
	}
}

// GetEntity allows us to get back any entity and
// their corresponding matching pattern from our LookupExtractor dict
func (ext *LookupExtractor) GetEntity(s string) map[string]interface{} {
	res := make(map[string]interface{})
	var (
		patternList []string
		re          *regexp.Regexp
	)
	for k, v := range ext.LookupTable {
		for _, pattern := range v {
			re = adjustPattern(pattern, ext.Flags)
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
// that match a pattern from our LookupExtractor dict
func (ext *LookupExtractor) GetSentences(slice []string) (res []string) {
	for _, v := range ext.LookupTable {
		for _, pattern := range v {
			for _, val := range slice {
				re := adjustPattern(pattern, ext.Flags)
				if match := re.FindString(val); match != "" {
					res = append(res, val)
				}
			}
		}
	}
	return
}

// adjustPattern adjusts the pattern entered with
// flags selected and returns the new pattern
func adjustPattern(pattern string, flags map[RegexFlag]bool) *regexp.Regexp {
	var opts string
	for flag, activated := range flags {
		if flag == IGNORECASE && activated {
			opts += "(?i)"
		}
		if flag == MULTILINE && activated {
			opts += "(?m)"
		}
		if flag == MATCHNEWLINE && activated {
			opts += "(?s)"
		}
		if flag == UNGREEDY && activated {
			opts += "(?U)"
		}
	}
	return regexp.MustCompile(opts + pattern)
}
