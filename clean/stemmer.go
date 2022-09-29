package clean

import (
	_ "embed"
	"encoding/json"
	"github.com/mx79/go-nlp/base"
	"github.com/mx79/go-nlp/utils"
	"log"
)

//go:embed ressources/stemmer.json
var stemmBytes []byte

// stemms is the map containing stemms for any language
var stemms = loadStemm(stemmBytes)

// Stemmer object
// In linguistics, stemming is a process of transforming inflections into their radical or root.
// The root of a word is the part of the word that remains after removing its prefix and suffix,
// namely its stem.
type Stemmer struct {
	Lang base.Lang
	Dict base.StemmDict
}

// loadStemm loads the map that contains the stemms in many languages
func loadStemm(b []byte) base.Stemms {
	var st base.Stemms
	utils.Check(json.Unmarshal(b, &st))
	return st
}

// NewStemmer instantiates a new Stemmer object
func NewStemmer(lang base.Lang) *Stemmer {
	return &Stemmer{
		Lang: lang,
		Dict: stemmDict(lang, stemms),
	}
}

// stemmDict retrieves a map of stemms for a language
func stemmDict(lang base.Lang, st base.Stemms) map[string]string {
	if _, ok := st[lang]; !ok {
		log.Fatal(base.LangError)
	}
	return st[lang]
}

// Stem is the method that is stemming every word in the input sentence
func (stm *Stemmer) Stem(s string) string {
	var sent string
	for _, word := range Tokenize(s, false) {
		if !utils.MapContains(stm.Dict, word) {
			sent += word + " "
		} else {
			sent += stm.Dict[word] + " "
		}
	}
	return sent
}
