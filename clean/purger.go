package clean

import (
	"github.com/mx79/go-nlp/clean/base"
)

const lemmStemmError = "You cannot use lemmatizing and stemming on the same text, this will leeds to strange results"

// TextPurger object
type TextPurger struct {
	Language    base.Lang
	Stopwords   *Stopwords
	Lemmatizer  *Lemmatizer
	Stemmer     *Stemmer
	NoStopword  bool
	Lemmatizing bool
	Stemming    bool
	NoPunct     bool
	NoAccent    bool
	Lowercase   bool
}

// NewTextPurger
func NewTextPurger(lang base.Lang, noStopword bool, lemmatizing bool, stemming bool, noPunct bool, noAccent bool, lowercase bool) *TextPurger {
	if lemmatizing && stemming {
		panic(lemmStemmError)
	}
	return &TextPurger{
		Language:    lang,
		Stopwords:   NewStopwords(lang),
		Lemmatizer:  NewLemmatizer(lang),
		Stemmer:     NewStemmer(lang),
		NoStopword:  noStopword,
		Lemmatizing: lemmatizing,
		Stemming:    stemming,
		NoPunct:     noPunct,
		NoAccent:    noAccent,
		Lowercase:   lowercase,
	}
}

// Purge The function that allows to clean a given text in depth
// by applying several layers of treatment
// It returns the sentence based on boolean values
func (p *TextPurger) Purge(s string) string {
	if p.NoStopword {
		s = p.Stopwords.Stop(s)
	}
	if p.Lemmatizing {
		s = p.Lemmatizer.Lemm(s)
	}
	if p.Stemming {
		s = p.Stemmer.Stem(s)
	}
	if p.NoPunct {
		s = RemovePunctuation(s)
	}
	if p.NoAccent {
		s = RemoveAccent(s)
	}
	if p.Lowercase {
		s = Lower(s)
	}
	return s
}
