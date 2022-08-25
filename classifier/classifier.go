package classifier

import (
	"encoding/json"
	"github.com/mx79/go-nlp/nlu"
	"io"
	"os"
)

// IntentClassifier
type IntentClassifier struct {
	SentPerIntent []int
	Vocab         []string
	Classes       []string
	Encoder       *nlu.LabelEncoder
	MatrixX       [][]float64
	MatrixY       []int
}

// NewIntentClassifier
func NewIntentClassifier(dataPath string) *IntentClassifier {
	// Opening json file
	jsonFile, err := os.Open(dataPath)
	if err != nil {
		panic(err)
	}
	defer jsonFile.Close()
	// Reading JSON file
	byteValue, _ := io.ReadAll(jsonFile)
	var results map[string][]string
	err = json.Unmarshal(byteValue, &results)
	if err != nil {
		panic(err)
	}
	// x, y, classes, vocab
	x, y, classes, vocab := nlu.BOW(results)
	// FitTransform like a LabelEncoder and get an encoded stopwords in return
	encoder := nlu.NewlabelEncoder()
	encodedY := encoder.Encode(y)
	// Creating slice to see how many sentence there is per intent
	//var sentPerIntent []int
	//for _, intent := range classes {
	//	for key, v := range results {
	//		if intent == key {
	//			sentPerIntent = append(sentPerIntent, len(v))
	//		}
	//	}
	//}
	/*
		// Resizing array to tranqform x in a matrix at the next step
		resizedArray := func(ndArray [][]float64) []float64 {
			var resizedArray []float64
			for _, array := range ndArray {
				for _, nb := range array {
					resizedArray = append(resizedArray, nb)
				}
			}
			return resizedArray
		}(x)
		// Transforming x and y into matrix
		X := mat.NewDense(len(x), len(x[0]), resizedArray)
		Y := mat.NewDense(len(y), 1, encodedY)
	*/
	// Returning an IntentClassifier
	return &IntentClassifier{
		Vocab:   vocab,
		Classes: classes,
		Encoder: encoder,
		MatrixX: x,
		MatrixY: encodedY,
	}
}

// bowSimilarity
func (cls *IntentClassifier) bowSimilarity(bow []float64) []float64 {
	// Variables
	var (
		count           = 0
		test            = 0
		i               = 1
		similarityCount int
		res             []float64
	)
	// Loop checking similarity between bow to test and bows of the train Matrix
	for _, b := range cls.MatrixX {
		nb := cls.SentPerIntent[test]
		for idx := range b {
			// Similarity between bows
			if b[idx] != 0 && bow[idx] != 0 {
				count++
			}
		}
		similarityCount += count
		count = 0
		if i == nb {
			test++
			// fmt.Printf("simi: %v, nb: %v\n", similarityCount, nb)
			res = append(res, float64(similarityCount/nb))
			similarityCount = 0
			i = 0
		}
		i++
	}
	return res
}

// GetIntent
func (cls *IntentClassifier) GetIntent(s string, boolSplitOnConj bool) string {
	bow := nlu.SentenceBOW(s, cls.Vocab)
	res, _ := idxOfMaxVal(cls.bowSimilarity(bow))
	pred := cls.Encoder.DecodeRes(res)
	return pred
}

// idxOfMaxVal
func idxOfMaxVal(slice []float64) (int, float64) {
	var test = 0.0
	var maxIndex int
	for idx, val := range slice {
		if test < val {
			test = val
			maxIndex = idx
		}
	}
	return maxIndex, test
}
