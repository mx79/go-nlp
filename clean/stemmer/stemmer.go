package stemmer

import (
	_ "embed"
	"encoding/json"
	"github.com/mx79/go-nlp/clean"
	"github.com/mx79/go-nlp/clean/base"
	"github.com/mx79/go-nlp/utils"
)

//go:embed stemmer.json
var stemmBytes []byte

// stemms
var stemms = loadStemm(stemmBytes)

// Stemmer
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
	for _, word := range clean.Tokenize(s) {
		if !utils.MapContains(stm.Dict, word) {
			sent += word
		} else {
			sent += stm.Dict[word]
		}
		sent += " "
	}
	return sent
}
