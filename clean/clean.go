package clean

import (
	"bufio"
	"github.com/aaaton/golem/v4"
	"github.com/aaaton/golem/v4/dicts/fr"
	"github.com/mx79/go-nlp/utils"
	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
	"log"
	"os"
	"path"
	"strings"
	"unicode"
)

// lemmatizer instance
var lemmatizer, _ = golem.New(fr.New())

// PurgeText The function that allows to clean a given text in depth by applying several layers of treatment
// return: The sentence or word list based on boolean values
func PurgeText(s string, noPunctuation bool, lemmatize bool, lowercase bool, noAccent bool, noStopword bool, noLowTfidfWord bool) string {
	if noPunctuation {
		s = RemovePunctuation(s)
	}
	if lemmatize {
		s = Lemmatize(s)
	}
	if lowercase {
		s = Lower(s)
	}
	if noAccent {
		s = RemoveAccent(s)
	}
	if noStopword {
		s = Stopword(s)
	}
	//if noLowTfidfWord {
	//	s = LowTfidfWord(s)
	//}
	return s
}

// Tokenize The function that separates the words of a sentence to facilitate their study in NLP
func Tokenize(s string) []string {
	return strings.Split(strings.Trim(s, " "), " ")
}

// slash trim the left side of a text containing ', useful to lem "j'ai" -> "ai" = "avoir"
func slash(word string) string {
	if strings.Contains(word, "'") {
		return strings.Split(word, "'")[1]
	}
	return word
}

// Lemmatize The function that lemmatizes the words of a sentence or a word
func Lemmatize(s string) string {
	var sent string
	for _, word := range Tokenize(s) {
		word = slash(word)
		sent += lemmatizer.Lemma(word) + " "
	}
	return sent
}

// RemovePunctuation The function that allows to remove punctuation in a sentence
func RemovePunctuation(s string) string {
	punctuation := "!@#$%^&*()[]_+<>?:.,;"
	for _, c := range punctuation {
		s = strings.Replace(s, string(c), "", -1)
	}
	return s
}

// RemoveAccent The function that allows you to remove the accents in a sentence
func RemoveAccent(s string) string {
	t := transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC)
	output, _, e := transform.String(t, s)
	if e != nil {
		panic(e)
	}
	return output
}

// stopwordList The function that retrieves a list of stopwords, here in French
func stopwordList() []string {
	var stopwords []string
	dir := path.Join("./stopwords", "fr.txt")
	data, err := os.Open(dir)
	if err != nil {
		log.Fatalf("Error when opening file: %s", err)
	}
	defer data.Close()
	fileScanner := bufio.NewScanner(data)
	for fileScanner.Scan() {
		stopwords = append(stopwords, fileScanner.Text())
	}
	if err = fileScanner.Err(); err != nil {
		log.Fatalf("Error while reading file: %s", err)
	}
	return stopwords
}

// Stopword The function that removes the stopwords contained in the input sentence
func Stopword(s string) string {
	var finalSent []string
	stopList := stopwordList()
	for _, word := range Tokenize(s) {
		if !utils.SliceContains(stopList, word) {
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
//// LowTfidfWord
//func LowTfidfWord(s string) string {
//	var finalSent []string
//	stopList := lowTfidfWordList()
//	for _, word := range Tokenize(s) {
//		if !utils.SliceContains(stopList, word) {
//			finalSent = append(finalSent, word)
//		}
//	}
//	return utils.ListToStr(finalSent)
//}

// Lower The function that put words of a sentence in lowercase mode
func Lower(s string) string {
	return strings.ToLower(s)
}
