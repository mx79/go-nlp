package distance

import (
	"fmt"
	"log"
	"strconv"
	"testing"
)

func AssertPass(t *testing.T, got, want interface{}) {
	if got != want {
		t.Fatalf("=> Got: %v\n=> Want: %v", got, want)
	}
}

func Round(x float64) float64 {
	if x == 0.0 {
		return x
	}

	trimmed, err := strconv.ParseFloat(fmt.Sprintf("%v", x)[:4], 64)
	if err != nil {
		log.Fatal(err)
	}

	return trimmed
}

var valWordErrorRate = []struct {
	sentenceA string
	sentenceB string
	want      float64
}{
	{"A word is missing", "A is missing", 0.33},
	{"There is no diff between this two sentences", "There is diff between this two sentences", 0.57},
	{"I see a difference here", "But not here", 0.33},
}

func TestWordErrorRate(t *testing.T) {
	for _, test := range valWordErrorRate {
		got := Round(WordErrorRate(test.sentenceA, test.sentenceB))
		AssertPass(t, got, test.want)
	}
}
