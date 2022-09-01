package clean

//// Purger
//type Purger struct {
//	Language       string
//	LemmInit       bool
//	Stopwords      []string
//	Lemmatizer     map[string]string
//	Stemmer        map[string]string
//	NoPunct        bool
//	Lemmatize      bool
//	Lowercase      bool
//	NoAccent       bool
//	NoStopword     bool
//	NoLowTfidfWord bool
//}
//
//// NewPurger
//func NewPurger(lang string, noPunct bool, lemmatize bool, lowercase bool, noAccent bool, noStopword bool, noLowTfidfWord bool) *Purger {
//	return &Purger{
//		Language:       lang,
//		Stopwords:      stopwordList(lang),
//		Lemmatizer:     lemmDict(lang),
//		Stemmer:        stemmDict(lang),
//		NoPunct:        noPunct,
//		Lemmatize:      lemmatize,
//		Lowercase:      lowercase,
//		NoAccent:       noAccent,
//		NoStopword:     noStopword,
//		NoLowTfidfWord: noLowTfidfWord,
//	}
//}
//
//// PurgeText The function that allows to clean a given text in depth by applying several layers of treatment
//// return: The sentence or word list based on boolean values
//func (p *Purger) PurgeText(s string) string {
//	if p.NoPunct {
//		s = removePunctuation(s)
//	}
//	if p.Lemmatize {
//		s = p.lemmatize(s)
//	}
//	if p.Lowercase {
//		s = lower(s)
//	}
//	if p.NoAccent {
//		s = removeAccent(s)
//	}
//	if p.NoStopword {
//		s = stopword(s, p.Stopwords)
//	}
//	//if noLowTfidfWord {
//	//	s = LowTfidfWord(s)
//	//}
//	return s
//}
