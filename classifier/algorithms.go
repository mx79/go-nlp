package classifier

import (
	"github.com/mx79/go-nlp/clean"
	"github.com/mx79/go-nlp/utils"
	"math"
	"strings"
)

// cleanAndTokenize is a local utility cleaning func implementing the clean package of go-nlp
func cleanAndTokenize(d string) (cleanedSent []string) {
	pur := clean.NewTextPurger("fr", true, true, true, true, true)
	cleanedSent = clean.Tokenize(pur.Purge(d), true)

	return
}

// BOW implements the bag of words algorithm.
func BOW(dataset map[string][]string) (x [][]float64, y []string, classes []string, vocab []string) {
	// Creating string slice and WordSet of the dataset
	var words, docX, docY []string
	for intent, patternsList := range dataset {
		for _, sent := range patternsList {
			tokens := cleanAndTokenize(sent)
			for _, token := range tokens {
				words = append(words, token)
			}
			docX = append(docX, sent)
			docY = append(docY, intent)
		}
	}
	classes = utils.SortedSet(docY)
	vocab = utils.SortedSet(words)
	if utils.SliceContains(vocab, "") {
		vocab = utils.SliceDeleteItem(vocab, "")
	}

	// Bag of words
	for idx, doc := range docX {
		var bow []float64
		text := cleanAndTokenize(doc)
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

	return
}

// SentenceBOW applies BOW algorithm but only on one sentence.
//
// This is used to get a vector from the sentence in order, most of the time,
// to process it in a classifier for example
func SentenceBOW(sentence string, vocab []string) (bow []float64) {
	text := cleanAndTokenize(sentence)
	for _, word := range vocab {
		if utils.SliceContains(text, word) {
			count := float64(strings.Count(utils.ListToStr(text), word))
			bow = append(bow, count)
		} else {
			bow = append(bow, 0)
		}
	}

	return
}

// TfidfVectorizer implements the TFIDF score calculation.
//
// It permits to understand the importance of a word in a sentence or a corpus of text.
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

	return tfidf
}
