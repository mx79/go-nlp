package lemmatizer

import (
	_ "embed"
	"encoding/json"
	"github.com/mx79/go-nlp/clean"
	"github.com/mx79/go-nlp/clean/base"
	"github.com/mx79/go-nlp/utils"
	"strings"
)

//go:embed lemmatizer.json
var lemmBytes []byte

// lemms
var lemms = loadLemm(lemmBytes)

// Lemmatizer
type Lemmatizer struct {
	Language base.Lang
	Dict     base.LemmDict
}

// loadLemm
func loadLemm(b []byte) base.Lemms {
	var l base.Lemms
	utils.Check(json.Unmarshal(b, &l))
	return l
}

// NewLemmatizer
func NewLemmatizer(lang base.Lang) *Lemmatizer {
	return &Lemmatizer{
		Language: lang,
		Dict:     lemmDict(lang, lemms),
	}
}

// lemmDict
func lemmDict(lang base.Lang, l base.Lemms) map[string]string {
	if _, ok := l[lang]; !ok {
		panic(base.LangError)
	}
	return l[lang]
}

// slash trim the left side of a text containing ', useful to lem "j'ai" -> "ai" = "avoir"
func slash(word string) string {
	if strings.Contains(word, "'") {
		return strings.Split(word, "'")[1]
	}
	return word
}

// Lemm
func (stm *Lemmatizer) Lemm(s string) string {
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
