package clean

import (
	_ "embed"
	"encoding/json"
	"github.com/mx79/go-nlp/base"
	"github.com/mx79/go-nlp/utils"
	"strings"
)

//go:embed ressources/lemmatizer.json
var lemmBytes []byte

// lemms is the map containing lemms for any language
var lemms = loadLemm(lemmBytes)

// Lemmatizer object
// Lemmatization refers to a lexical treatment of a text
// in order to classify it in an index or to analyse it.
// This treatment consists in applying to the occurrences of lexemes subject
// to inflection a coding referring to their common lexical entry,
// which is referred to as a lemma.
type Lemmatizer struct {
	Language base.Lang
	Dict     base.LemmDict
}

// loadLemm loads the map that contains the lemms in many languages
func loadLemm(b []byte) base.Lemms {
	var l base.Lemms
	utils.Check(json.Unmarshal(b, &l))
	return l
}

// NewLemmatizer instantiates a new Lemmatizer object
func NewLemmatizer(lang base.Lang) *Lemmatizer {
	return &Lemmatizer{
		Language: lang,
		Dict:     lemmDict(lang, lemms),
	}
}

// lemmDict retrieves a map of lemms for a language
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

// Lemm is the method that is lemmatizing every word in the input sentence
func (stm *Lemmatizer) Lemm(s string) string {
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
