package clean

import "testing"

var (
	valFrStem = []struct {
		sentence string
		want     string
	}{
		{"Je fais à les test là sur les mots", "Je fais à le test là sur le mot"},
		{"On voit si les stopwords sont enlevés ou non", "On voit si le stopwords sont enlevés ou non"},
	}
	valEnStem = []struct {
		sentence string
		want     string
	}{
		{"I think I want to eliminate some words in the text here, let's see which one are removed",
			"I think I want to elimin some word in the text here, let's see which one are remov"},
		{"I'm doing some tests", "I'm do some test"},
	}
)

func TestStemmer_Stem(t *testing.T) {
	frSt := NewStemmer(FR)
	for _, test := range valFrStem {
		got := frSt.Stem(test.sentence)
		AssertPass(t, got, test.want)
	}

	enSt := NewStemmer(EN)
	for _, test := range valEnStem {
		got := enSt.Stem(test.sentence)
		AssertPass(t, got, test.want)
	}
}
