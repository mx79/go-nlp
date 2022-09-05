package clean

import (
	_ "embed"
	"encoding/json"
	"github.com/mx79/go-nlp/clean/base"
	"github.com/mx79/go-nlp/utils"
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

// loadStemm
func loadStemm(b []byte) base.Stemms {
	var st base.Stemms
	utils.Check(json.Unmarshal(b, &st))
	return st
}

// NewStemmer
func NewStemmer(lang base.Lang) *Stemmer {
	return &Stemmer{
		Lang: lang,
		Dict: stemmDict(lang, stemms),
	}
}

// stemmDict
func stemmDict(lang base.Lang, st base.Stemms) map[string]string {
	if _, ok := st[lang]; !ok {
		panic(base.LangError)
	}
	return st[lang]
}

// Stem
func (stm *Stemmer) Stem(s string) string {
	var sent string
	for _, word := range Tokenize(s) {
		if !utils.MapContains(stm.Dict, word) {
			sent += word
		} else {
			sent += stm.Dict[word]
		}
		sent += " "
	}
	return sent
}
