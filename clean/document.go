package clean

import (
	"github.com/mx79/go-nlp/utils"
	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
	"strings"
	"unicode"
)

var stopwords = map[string][]string{}

var lemmatizer = map[string][]string{}

var stemmer = map[string][]string{}

// Doc
type Doc struct {
	string
	*Purger
}

// Purger
type Purger struct {
	Language       string
	LemmInit       bool
	Stopwords      []string
	Lemmatizer     []string
	Stemmer        []string
	NoPunct        bool
	Lemmatize      bool
	Lowercase      bool
	NoAccent       bool
	NoStopword     bool
	NoLowTfidfWord bool
}

// NewPurger
func NewPurger(lang string, noPunct bool, lemmatize bool, lowercase bool, noAccent bool, noStopword bool, noLowTfidfWord bool) *Purger {
	return &Purger{
		Language:       lang,
		Stopwords:      stopwordList(lang),
		Lemmatizer:     lemmatizerList(lang),
		Stemmer:        stemmerList(lang),
		NoPunct:        noPunct,
		Lemmatize:      lemmatize,
		Lowercase:      lowercase,
		NoAccent:       noAccent,
		NoStopword:     noStopword,
		NoLowTfidfWord: noLowTfidfWord,
	}
}

// PurgeText The function that allows to clean a given text in depth by applying several layers of treatment
// return: The sentence or word list based on boolean values
func (p *Purger) PurgeText(s string) string {
	if p.NoPunct {
		s = removePunctuation(s)
	}
	if p.Lemmatize {
		s = p.lemmatize(s)
	}
	if p.Lowercase {
		s = lower(s)
	}
	if p.NoAccent {
		s = removeAccent(s)
	}
	if p.NoStopword {
		s = stopword(s, p.Stopwords)
	}
	//if noLowTfidfWord {
	//	s = LowTfidfWord(s)
	//}
	return s
}

// slash trim the left side of a text containing ', useful to lem "j'ai" -> "ai" = "avoir"
func slash(word string) string {
	if strings.Contains(word, "'") {
		return strings.Split(word, "'")[1]
	}
	return word
}

// lemmatize The function that lemmatizes the words of a sentence or a word
func (p *Purger) lemmatize(s string) string {
	if p.LemmInit {
		var sent string
		for _, word := range Tokenize(s) {
			sent += p.Lemmatizer.Lemma(slash(word)) + " "
		}
		return sent
	} else {
		panic("Purger cannot use func lemmatize if attribute LemmInit is not set to true")
	}
}

// removePunctuation The function that allows to remove punctuation in a sentence
func removePunctuation(s string) string {
	punctuation := "!@#$%^&*()[]_+<>?:.,;"
	for _, c := range punctuation {
		s = strings.Replace(s, string(c), "", -1)
	}
	return s
}

// removeAccent The function that allows you to remove the accents in a sentence
func removeAccent(s string) string {
	t := transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC)
	output, _, err := transform.String(t, s)
	if err != nil {
		panic(err)
	}
	return output
}

// stopwordList The function that retrieves a list of stopwords
func stopwordList(lang string) []string {
	if _, ok := stopwords[lang]; !ok {
		panic("lang variable value is not recognized, check the lang code you want, " +
			"for example: 'de' for German, 'en' for English")
	}
	return stopwords[lang]
}

// stopword The function that removes the stopwords contained in the input sentence
func stopword(s string, stopwords []string) string {
	var finalSent []string
	for _, word := range Tokenize(s) {
		if !utils.SliceContains(stopwords, word) {
			finalSent = append(finalSent, word)
		}
	}
	return utils.ListToStr(finalSent)
}

//// stopwordList The function that retrieves a list of stopwords, here in French
//func lowTfidfWordList() []string {
//	var lowTfidfWord []string
//	dir := path.Join(utils.ChuchoDir, "stopwords", "lowTfidfWord.txt")
//	stopwords, e2 := os.Open(dir)
//	if e2 != nil {
//		log.Fatalf("Error when opening file: %s", e2)
//	}
//	defer stopwords.Close()
//	fileScanner := bufio.NewScanner(stopwords)
//	for fileScanner.Scan() {
//		lowTfidfWord = append(lowTfidfWord, fileScanner.Text())
//	}
//	if e3 := fileScanner.Err(); e3 != nil {
//		log.Fatalf("Error while reading file: %s", e3)
//	}
//	// stopwords.Close()
//	return lowTfidfWord
//}
//
//// lowTfidfWord
//func lowTfidfWord(s string) string {
//	var finalSent []string
//	stopList := lowTfidfWordList()
//	for _, word := range Tokenize(s) {
//		if !utils.SliceContains(stopList, word) {
//			finalSent = append(finalSent, word)
//		}
//	}
//	return utils.ListToStr(finalSent)
//}

// lower The function that put words of a sentence in lowercase mode
func lower(s string) string {
	return strings.ToLower(s)
}

// Tokenize The function that separates the words of a sentence to facilitate their study in NLP
func Tokenize(s string) []string {
	return strings.Split(strings.Trim(s, " "), " ")
}
