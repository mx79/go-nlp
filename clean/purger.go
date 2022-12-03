package clean

import "strings"

// TextPurger object with boolean attributes to specify
// which clean function we want to use or not.
type TextPurger struct {
	Language    Lang
	Stopwords   *Stopwords
	Stemmer     *Stemmer
	NoStopword  bool
	Lemmatizing bool
	Stemming    bool
	NoPunct     bool
	NoAccent    bool
	Lowercase   bool
}

// NewTextPurger instantiates a new TextPurger object
func NewTextPurger(lang Lang, noStopword bool, stemming bool, noPunct bool, noAccent bool, lowercase bool) *TextPurger {
	return &TextPurger{
		Language:   lang,
		Stopwords:  NewStopwords(lang),
		Stemmer:    NewStemmer(lang),
		NoStopword: noStopword,
		Stemming:   stemming,
		NoPunct:    noPunct,
		NoAccent:   noAccent,
		Lowercase:  lowercase,
	}
}

// Purge allows to clean a given text in depth
// by applying several layers of treatment.
//
// It returns the sentence based on boolean values
func (p *TextPurger) Purge(s string) string {
	if p.NoPunct {
		s = RemovePunctuation(s)
	}
	if p.NoAccent {
		s = RemoveAccent(s)
	}
	if p.Lowercase {
		s = strings.ToLower(s)
	}
	if p.NoStopword {
		s = p.Stopwords.Stop(s)
	}
	if p.Stemming {
		s = p.Stemmer.Stem(s)
	}

	return s
}
