package classifier

import (
	"encoding/json"
	"io"
	"log"
	"os"
)

// IntentClassifier object
//
// This object make possible the detection of an intention in an input sentence
type IntentClassifier struct {
	SentPerIntent []int
	Vocab         []string
	Classes       []string
	Encoder       *LabelEncoder
	MatrixX       [][]float64
	MatrixY       []int
}

// NewIntentClassifier instantiates a new IntentClassifier object
func NewIntentClassifier(dataPath string) *IntentClassifier {
	// Opening json file
	jsonFile, err := os.Open(dataPath)
	if err != nil {
		log.Fatal(err)
	}
	defer jsonFile.Close()

	// Reading JSON file
	byteValue, _ := io.ReadAll(jsonFile)
	var results map[string][]string
	err = json.Unmarshal(byteValue, &results)
	if err != nil {
		log.Fatal(err)
	}

	// x, y, classes, vocab
	x, y, classes, vocab := BOW(results)

	// FitTransform like a LabelEncoder and get an encoded stopwords in return
	encoder := NewLabelEncoder()
	encoder.Encode(y)

	//Creating slice to see how many sentence there is per intent
	sentPerIntent := func() (sentPerIntent []int) {
		for _, intent := range classes {
			for key, v := range results {
				if intent == key {
					sentPerIntent = append(sentPerIntent, len(v))
				}
			}
		}

		return
	}()

	//// Resizing array to tranqform x in a matrix at the next step
	//resizedArray := func() (arr []float64) {
	//	for _, array := range x {
	//		for _, nb := range array {
	//			arr = append(arr, nb)
	//		}
	//	}
	//
	//	return
	//}()

	//// Transforming x and y into matrix
	//X := mat.NewDense(len(x), len(x[0]), resizedArray)
	//Y := mat.NewDense(len(y), 1, encodedY)

	return &IntentClassifier{
		SentPerIntent: sentPerIntent,
		Vocab:         vocab,
		Classes:       classes,
		Encoder:       encoder,
		MatrixX:       x,
		MatrixY:       encoder.EncodedData,
	}
}

// GetIntent takes a sentence as an entry and returns a map of intent and probability score for each one.
func (cls *IntentClassifier) GetIntent(s string) map[string]float64 {
	res := make(map[string]float64)
	bow := SentenceBOW(s, cls.Vocab)
	pred := cls.Encoder.DecodeRes(idxOfMaxVal(bow))
	res[pred] = 0.0 // TODO: Score

	return res
}

// idxOfMaxVal
func idxOfMaxVal(slice []float64) (maxIndex int) {
	var test = 0.0
	for idx, val := range slice {
		if test < val {
			test = val
			maxIndex = idx
		}
	}

	return
}
