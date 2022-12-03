package clean

import "testing"

func AssertPass(t *testing.T, got, want interface{}) {
	if got != want {
		t.Fatalf("=> Got: %v\n=> Want: %v", got, want)
	}
}

var (
	valFrStop = []struct {
		sentence string
		want     string
	}{
		{"Je fais à les test là sur les mots", "Je fais à test mots"},
		{"On voit si les stopwords sont enlevés ou non", "On voit stopwords enlevés non"},
	}
	valEnStop = []struct {
		sentence string
		want     string
	}{
		{"I think I want to eliminate some words in the text here, let's see which one are removed",
			"I think I want eliminate words text here, see one removed"},
		{"I'm doing some tests", "I'm tests"},
	}
)

func TestStopwords_Stop(t *testing.T) {
	frSt := NewStopwords(FR)
	for _, test := range valFrStop {
		got := frSt.Stop(test.sentence)
		AssertPass(t, got, test.want)
	}

	enSt := NewStopwords(EN)
	for _, test := range valEnStop {
		got := enSt.Stop(test.sentence)
		AssertPass(t, got, test.want)
	}
}
