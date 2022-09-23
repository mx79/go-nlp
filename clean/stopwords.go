package clean

import (
	_ "embed"
	"encoding/json"
	"github.com/mx79/go-nlp/base"
	"github.com/mx79/go-nlp/utils"
)

//go:embed ressources/stopwords.json
var stopBytes []byte

// stopwords is the map containing stopwords in many languages
var stopwords = loadStop(stopBytes)

// Stopwords object
// In information retrieval, a stopword is a word that is so common
// that there is no need to index it or use it in a search.
type Stopwords struct {
	Language base.Lang
	List     []string
}

// loadStop loads the map that contains the stopwords in many languages
func loadStop(b []byte) base.GlobalStopwords {
	var s base.GlobalStopwords
	utils.Check(json.Unmarshal(b, &s))
	return s
}

// NewStopwords instantiates a new Stopwords object
func NewStopwords(lang base.Lang) *Stopwords {
	return &Stopwords{
		Language: lang,
		List:     stopwordList(lang, stopwords),
	}
}

// stopwordList retrieves a list of stopwords for a language
func stopwordList(lang base.Lang, s base.GlobalStopwords) []string {
	if _, ok := s[lang]; !ok {
		panic(base.LangError)
	}
	return s[lang]
}

// Stop is the method that removes the stopwords contained in the input sentence
func (stp *Stopwords) Stop(s string) string {
	var sent string
	for _, word := range Tokenize(s) {
		if !utils.SliceContains(stp.List, word) {
			sent += word + " "
		}
	}
	return sent
}
