package classifier

import (
	"github.com/mx79/go-nlp/utils"
	"log"
	"sort"
)

const missingSchemeError = "Cannot decode labels which are not encoded, use Encode() method first"

// LabelEncoder is inspired from Python sklearn library.
//
// It allows to create a slice of number based on the number of input labels.
//
// It makes easier the use of mathematical functions on number than on string.
//
// We can then reverse the encoding process to get back our initial slice of string.
type LabelEncoder struct {
	EncodingScheme map[string]int
	EncodedData    []int
}

// NewLabelEncoder instantiates a new LabelEncoder
func NewLabelEncoder() *LabelEncoder {
	return &LabelEncoder{}
}

// Encode FitTransform like a label encoder from sklearn.
//
// It will store a map of the encoding scheme in the attribute encodedData, also the encodingScheme.
func (enc *LabelEncoder) Encode(stringSlice []string) {
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
	enc.EncodedData = encodedSlice
}

// Decode is the reverse action of the encoding function.
//
// It will return the initial slice of string provided before encoding.
func (enc *LabelEncoder) Decode() (stringSlice []string) {
	if enc.EncodingScheme != nil {
		for _, val := range enc.EncodedData {
			for k, v := range enc.EncodingScheme {
				if v == val {
					stringSlice = append(stringSlice, k)
				}
			}
		}
	} else {
		log.Fatal(missingSchemeError)
	}

	return
}

// DecodeRes will return the label of the input int,
// according to the encodingScheme our object have.
func (enc *LabelEncoder) DecodeRes(res int) (key string) {
	if enc.EncodingScheme != nil {
		for k, v := range enc.EncodingScheme {
			if v == res {
				key = k
				break
			}
		}
	} else {
		log.Fatal(missingSchemeError)
	}

	return
}
