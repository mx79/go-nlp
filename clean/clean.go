package clean

import (
	"bufio"
	"fmt"
	"github.com/aaaton/golem/v4"
	"github.com/aaaton/golem/v4/dicts/de"
	"github.com/aaaton/golem/v4/dicts/en"
	"github.com/aaaton/golem/v4/dicts/es"
	"github.com/aaaton/golem/v4/dicts/fr"
	"github.com/aaaton/golem/v4/dicts/it"
	"github.com/aaaton/golem/v4/dicts/sv"
	"github.com/mx79/go-nlp/utils"
	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
	"log"
	"os"
	"path"
	"runtime"
	"strings"
	"unicode"
)

var (
	_, f, _, _ = runtime.Caller(0)
	dir        = path.Dir(f)
)

// Purger
type Purger struct {
	Language       string
	LemmInit       bool
	Lemmatizer     *golem.Lemmatizer
	Stopwords      []string
	NoPunct        bool
	Lemmatize      bool
	Lowercase      bool
	NoAccent       bool
	NoStopword     bool
	NoLowTfidfWord bool
}

// NewPurger
func NewPurger(lang string, initLemm bool, noPunct bool, lemmatize bool, lowercase bool, noAccent bool, noStopword bool, noLowTfidfWord bool) *Purger {
	var (
		langPack   golem.LanguagePack
		lemmatizer *golem.Lemmatizer
	)
	if initLemm {
		if lang != "fr" && lang != "en" && lang != "es" && lang != "it" && lang != "de" && lang != "sv" {
			panic("lang variable value is not recognized, use on of the following: " +
				"'fr', 'en', 'es, 'it', 'de', 'sv'")
		} else if lang == "fr" {
			langPack = fr.New()
		} else if lang == "en" {
			langPack = en.New()
		} else if lang == "es" {
			langPack = es.New()
		} else if lang == "it" {
			langPack = it.New()
		} else if lang == "de" {
			langPack = de.New()
		} else if lang == "sv" {
			langPack = sv.New()
		}
		lemmatizer, _ = golem.New(langPack)
	}
	return &Purger{
		Language:       lang,
		LemmInit:       initLemm,
		Lemmatizer:     lemmatizer,
		Stopwords:      stopwordList(lang),
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
	var stopwords []string
	if lang == "fr" {
		dir = path.Join(dir, "stopwords", "fr.txt")
	} else if lang == "en" {
		dir = path.Join(dir, "stopwords", "en.txt")
	} else if lang == "es" {
		dir = path.Join(dir, "stopwords", "es.txt")
	} else if lang == "it" {
		dir = path.Join(dir, "stopwords", "it.txt")
	} else if lang == "de" {
		dir = path.Join(dir, "stopwords", "de.txt")
	} else if lang == "sv" {
		dir = path.Join(dir, "stopwords", "sv.txt")
	}
	if lang == "None" {
		fmt.Println("No stopword is set")
	} else {
		data, err := os.Open(dir)
		if err != nil {
			panic("lang variable value for stopwords is not recognized, use on of the following: " +
				"'fr', 'en', 'es, 'it', 'de', 'sv'")
		}
		defer data.Close()
		fileScanner := bufio.NewScanner(data)
		for fileScanner.Scan() {
			stopwords = append(stopwords, fileScanner.Text())
		}
		if err := fileScanner.Err(); err != nil {
			log.Fatalf("Error while reading file: %s", err)
		}
	}
	return stopwords
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
