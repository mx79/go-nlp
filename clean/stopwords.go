package clean

import (
	_ "embed"
	"encoding/json"
	"github.com/mx79/go-nlp/utils"
	"log"
	"strings"
)

//go:embed ressources/stopwords.json
var stopBytes []byte

// stopwords is the map containing stopwords in many languages
var stopwords = loadStop(stopBytes)

// Stopwords object
//
// In information retrieval, a stopword is a word that is so common
// that there is no need to index it or use it in a search.
type Stopwords struct {
	Language lang
	List     stopList
}

// loadStop loads the map that contains the stopwords in many languages
func loadStop(b []byte) globalStopwords {
	s := make(globalStopwords)
	err := json.Unmarshal(b, &s)
	if err != nil {
		log.Fatal(err)
	}

	return s
}

// NewStopwords instantiates a new Stopwords object
func NewStopwords(lang lang) *Stopwords {
	return &Stopwords{
		Language: lang,
		List:     stopwordList(lang, stopwords),
	}
}

// stopwordList retrieves a list of stopwords for a language
func stopwordList(lang lang, s globalStopwords) stopList {
	if _, ok := s[lang]; !ok {
		log.Fatal(LangError)
	}

	return s[lang]
}

// Stop is the method that removes the stopwords contained in the input sentence
func (stp *Stopwords) Stop(s string) string {
	var sent string
	for _, word := range Tokenize(s, false) {
		if !utils.SliceContains(stp.List, word) {
			sent += word + " "
		}
	}
	sent = strings.Trim(sent, " ")

	return strings.TrimSpace(sent)
}
