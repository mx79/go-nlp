package classifier

import (
	"encoding/json"
	"github.com/mx79/go-nlp/nlu"
	"github.com/mx79/go-nlp/utils"
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
	utils.Check(err)
	defer jsonFile.Close()
	// Reading JSON file
	byteValue, _ := io.ReadAll(jsonFile)
	var results map[string][]string
	utils.Check(json.Unmarshal(byteValue, &results))
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

/*
// GetIntent
func (cls *IntentClassifier) GetIntent(s string, boolSplitOnConj bool) map[string]float64 {
	res := make(map[string]float64)
	bow := nlu.SentenceBOW(s, cls.Vocab)
	pred := cls.Encoder.Decode(bow)
	r[pred] = 0.0 // TODO: Score
	return res
}
*/

/*
// idxOfMaxVal
func idxOfMaxVal(slice []float64) int {
	var test = 0.0
	var maxIndex int
	for idx, val := range slice {
		if test < val {
			test = val
			maxIndex = idx
		}
	}
	return maxIndex
}
*/
