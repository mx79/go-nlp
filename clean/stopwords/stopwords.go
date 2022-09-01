package stopwords

import (
	_ "embed"
	"encoding/json"
	"github.com/mx79/go-nlp/clean"
	"github.com/mx79/go-nlp/clean/base"
	"github.com/mx79/go-nlp/utils"
)

//go:embed stopwords.json
var stopBytes []byte

// stopwords
var stopwords = loadStop(stopBytes)

// Stopwords
type Stopwords struct {
	Language base.Lang
	List     []string
}

// loadStop
func loadStop(b []byte) base.GlobalStopwords {
	var s base.GlobalStopwords
	utils.Check(json.Unmarshal(b, &s))
	return s
}

// New
func New(lang base.Lang) *Stopwords {
	return &Stopwords{
		Language: lang,
		List:     stopwordList(lang, stopwords),
	}
}

// stopwordList The function that retrieves a list of stopwords
func stopwordList(lang base.Lang, s base.GlobalStopwords) []string {
	if _, ok := s[lang]; !ok {
		panic(base.LangError)
	}
	return s[lang]
}

// Stop The function that removes the stopwords contained in the input sentence
func (stp *Stopwords) Stop(s string) string {
	var sent string
	for _, word := range clean.Tokenize(s) {
		if !utils.SliceContains(stp.List, word) {
			sent += word + " "
		}
	}
	return sent
}
