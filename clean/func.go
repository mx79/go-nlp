package clean

import (
	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
	"strings"
	"unicode"
)

// lemmatize The function that lemmatizes the words of a sentence or a word
//func (p *Purger) lemmatize(s string) string {
//	if p.LemmInit {
//		var sent string
//		for _, word := range Tokenize(s) {
//			sent += p.Lemmatizer.Lemma(slash(word)) + " "
//		}
//		return sent
//	} else {
//		panic("Purger cannot use func lemmatize if attribute LemmInit is not set to true")
//	}
//}

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

// lower The function that put words of a sentence in lowercase mode
func lower(s string) string {
	return strings.ToLower(s)
}

// Tokenize The function that separates the words of a sentence to facilitate their study in NLP
func Tokenize(s string) []string {
	return strings.Split(strings.Trim(s, " "), " ")
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
