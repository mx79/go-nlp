package nlu

import (
	"github.com/mx79/go-nlp/clean"
	"github.com/mx79/go-nlp/utils"
	"math"
	"sort"
	"strings"
)

// CleanAndTokenize
func CleanAndTokenize(s string) []string {
	cleanedSent := clean.Tokenize(clean.PurgeText(s, true, true, true, true, true, true))
	return cleanedSent
}

// firstHandler
func firstHandler(dataset map[string][]string) ([]string, []string, []string, []string) {
	// Creating string slice and WordSet of the dataset
	var words, docX, docY, vocab []string
	for intent, patternsList := range dataset {
		for _, sent := range patternsList {
			tokens := CleanAndTokenize(sent)
			for _, token := range tokens {
				words = append(words, token)
			}
			docX = append(docX, sent)
			docY = append(docY, intent)
		}
	}
	classes := utils.SortedWordSet(docY)
	vocab = utils.SortedWordSet(words)
	if utils.SliceContains(vocab, "") {
		vocab = utils.SliceDeleteItem(vocab, "")
	}
	return docX, docY, classes, vocab
}

// BOW
func BOW(dataset map[string][]string) ([][]float64, []string, []string, []string) {
	// First extract from the input dataset
	docX, docY, classes, vocab := firstHandler(dataset)
	// Bag of words
	var (
		x [][]float64
		y []string
	)
	for idx, doc := range docX {
		var bow []float64
		text := CleanAndTokenize(doc)
		for _, word := range vocab {
			if utils.SliceContains(text, word) {
				count := float64(strings.Count(utils.ListToStr(text), word))
				bow = append(bow, count)
			} else {
				bow = append(bow, 0)
			}
		}
		x = append(x, bow)
		y = append(y, docY[idx])
	}
	return x, y, classes, vocab
}

// SentenceBOW
func SentenceBOW(sentence string, vocab []string) []float64 {
	var bow []float64
	text := CleanAndTokenize(sentence)
	for _, word := range vocab {
		if utils.SliceContains(text, word) {
			count := float64(strings.Count(utils.ListToStr(text), word))
			bow = append(bow, count)
		} else {
			bow = append(bow, 0)
		}
	}
	return bow
}

// TfidfVectorizer
func TfidfVectorizer(dataset map[string][]string) map[string]float64 {
	// N: the number of sentence in the dataset
	N := 0.0
	for _, v := range dataset {
		for range v {
			N++
		}
	}
	// tf: term frequency
	tf, _, _, vocab := BOW(dataset)
	// df: document frequency
	df := make([]float64, len(tf[0]))
	for _, bow := range tf {
		for idx, word := range bow {
			if word != 0 {
				df[idx] += word
			}
		}
	}
	// idf:  inverse document frequency
	idf := make([]float64, len(df))
	for idx, wordFreq := range df {
		idf[idx] = math.Log(N * wordFreq)
	}
	// tf-idf score for each sentence
	score := make([]float64, len(tf[0]))
	tfidf := make(map[string]float64)
	for _, bow := range tf {
		for i2 := range bow {
			score[i2] += bow[i2] * idf[i2]
		}
	}
	for idx := range vocab {
		tfidf[vocab[idx]] = score[idx] / df[idx]
	}
	// return the score
	return tfidf
}

//func WriteUselessWord(dataset map[string][]string) {
//	// Getting tfidf score from words in the input dataset
//	tfidf := TfidfVectorizer(dataset)
//	// Getting useless word according to their tf-idf score
//	useless := func(tfidf map[string]float64) []string {
//		var useless []string
//		for k, v := range tfidf {
//			if v < 7 {
//				useless = append(useless, k)
//			}
//		}
//		return useless
//	}(tfidf)
//	// Writing in the file
//	dir := path.Join()
//	file, err := os.OpenFile(dir, os.O_WRONLY, 0600)
//	defer file.Close()
//	utils.Check(err)
//	for _, word := range useless {
//		utils.Write(word+"\n", file)
//	}
//}

// LabelEncoder
type LabelEncoder struct {
	EncodingScheme map[string]int
}

// NewLabelEncoder
func NewlabelEncoder() *LabelEncoder {
	return &LabelEncoder{}
}

// Encode FitTransform like a label encoder from sklearn
// and get a map of the encoding scheme in return, also the encoded stopwords
func (enc *LabelEncoder) Encode(stringSlice []string) []int {
	sort.Strings(stringSlice)
	var encodedSlice []int
	encodingScheme := make(map[string]int)
	idx := 0
	for _, val := range stringSlice {
		_, keyExist := encodingScheme[val]
		if utils.SliceContains(stringSlice, val) && !keyExist {
			encodingScheme[val] = idx + 1
			idx++
		}
	}
	for _, intent := range stringSlice {
		for key := range encodingScheme {
			if key == intent {
				encodedSlice = append(encodedSlice, encodingScheme[key])
			}
		}
	}
	enc.EncodingScheme = encodingScheme
	return encodedSlice
}

// Decode
//func (enc *LabelEncoder) Decode(encodedSlice []int) []string {
//	var stringSlice []string
//	TODO: complete this func
//	return stringSlice
//}

// DecodeRes
func (enc *LabelEncoder) DecodeRes(res int) string {
	var key string
	if enc.EncodingScheme != nil {
		for k, v := range enc.EncodingScheme {
			if v == res {
				key = k
				break
			}
		}
	} else {
		panic("Cannot decode a non encoded stopwords, use Encode() method first")
	}
	return key
}
