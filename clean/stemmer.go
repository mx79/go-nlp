package clean

import (
	_ "embed"
	"encoding/json"
	"github.com/mx79/go-nlp/utils"
	"log"
	"strings"
)

//go:embed ressources/stemmer.json
var stemmBytes []byte

// stemms is the map containing stemms for any language
var stemms = loadStemm(stemmBytes)

// Stemmer object
//
// In linguistics, stemming is a process of transforming inflections into their radical or root.
// The root of a word is the part of the word that remains after removing its prefix and suffix,
// namely its stem.
type Stemmer struct {
	Lang Lang
	Dict StemmDict
}

// loadStemm loads the map that contains the stemms in many languages
func loadStemm(b []byte) (st Stemms) {
	err := json.Unmarshal(b, &st)
	if err != nil {
		log.Fatal(err)
	}

	return
}

// NewStemmer instantiates a new Stemmer object
func NewStemmer(lang Lang) *Stemmer {
	return &Stemmer{
		Lang: lang,
		Dict: stemmDict(lang, stemms),
	}
}

// stemmDict retrieves a map of stemms for a language
func stemmDict(lang Lang, st Stemms) StemmDict {
	if _, ok := st[lang]; !ok {
		log.Fatal(LangError)
	}

	return st[lang]
}

// Stem is the method that is stemming every word in the input sentence
//func (stm *Stemmer) Stem(s string) string {
//	var sent string
//	for _, word := range Tokenize(s, false) {
//		if !utils.MapContains(stm.Dict, word) {
//			sent += word + " "
//		} else {
//			sent += stm.Dict[word] + " "
//		}
//	}
//	return sent
//}

// Stem is the method that is stemming every word in the input sentence
func (stm *Stemmer) Stem(s string) string {
	for _, word := range Tokenize(s, false) {
		if utils.MapContains(stm.Dict, word) {
			s = strings.Replace(s, word, stm.Dict[word], -1)
		}
	}

	return s
}
