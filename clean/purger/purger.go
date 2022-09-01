package purger

import (
	"github.com/mx79/go-nlp/clean"
	"github.com/mx79/go-nlp/clean/base"
	"github.com/mx79/go-nlp/clean/lemmatizer"
	"github.com/mx79/go-nlp/clean/stemmer"
	"github.com/mx79/go-nlp/clean/stopwords"
)

// Purger
type Purger struct {
	Language    base.Lang
	Stopwords   *stopwords.Stopwords
	Lemmatizer  *lemmatizer.Lemmatizer
	Stemmer     *stemmer.Stemmer
	NoPunct     bool
	Lemmatizing bool
	Stemming    bool
	Lowercase   bool
	NoAccent    bool
	NoStopword  bool
}

// New
func New(lang base.Lang, noPunct bool, lemmatizing bool, stemming bool, lowercase bool, noAccent bool, noStopword bool) *Purger {
	if lemmatizing && stemming {
		panic("You cannot use lemmatizing and stemming in the same, this will leeds strange results")
	}
	return &Purger{
		Language:    lang,
		Stopwords:   stopwords.New(lang),
		Lemmatizer:  lemmatizer.New(lang),
		Stemmer:     stemmer.New(lang),
		NoPunct:     noPunct,
		Lemmatizing: lemmatizing,
		Stemming:    stemming,
		Lowercase:   lowercase,
		NoAccent:    noAccent,
		NoStopword:  noStopword,
	}
}

// CleanText The function that allows to clean a given text in depth by applying several layers of treatment
// return: The sentence or word list based on boolean values
func (p *Purger) CleanText(s string) string {
	if p.NoPunct {
		s = clean.RemovePunctuation(s)
	}
	if p.Lemmatizing {
		s = p.Lemmatizer.Lemm(s)
	}
	if p.Stemming {
		s = p.Stemmer.Stem(s)
	}
	if p.Lowercase {
		s = clean.Lower(s)
	}
	if p.NoAccent {
		s = clean.RemoveAccent(s)
	}
	if p.NoStopword {
		s = p.Stopwords.Stop(s)
	}
	return s
}
