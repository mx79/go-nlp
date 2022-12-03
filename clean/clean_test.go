package clean

import (
	"testing"
)

var valRemovePunctuation = []struct {
	sentence string
	want     string
}{
	{"dZLl! mep. :: ?dm, zm?", "dZLl mep  dm zm"},
	{"!@#$%^&*()[]_+<>?:.,;", ""},
	{"je fais, des ?tests !!est-ce que ça marche*", "je fais des tests est-ce que ça marche"},
}

func TestRemovePunctuation(t *testing.T) {
	for _, test := range valRemovePunctuation {
		got := RemovePunctuation(test.sentence)
		AssertPass(t, got, test.want)
	}
}

var valRemoveAccent = []struct {
	sentence string
	want     string
}{
	{"à des b c ü ê les", "a des b c u e les"},
	{"àà les é ù Ï", "aa les e u I"},
	{"ùùù ï î â aak", "uuu i i a aak"},
}

func TestRemoveAccent(t *testing.T) {
	for _, test := range valRemoveAccent {
		got := RemoveAccent(test.sentence)
		AssertPass(t, got, test.want)
	}
}

var valTokenize = []struct {
	sentence string
	want     []string
}{
	{"Je fais, les tests sur le tokenizer en place !", []string{"Je", "fais", ",", "les", "tests", "sur", "le", "tokenizer", "en", "place", "!"}},
	{"I'm doing some tests with the tokenizer here! See?", []string{"I", "'", "m", "doing", "some", "tests", "with", "the", "tokenizer", "here", "!", "See", "?"}},
	{"A last: test to see if this, is working well??", []string{"A", "last", ":", "test", "to", "see", "if", "this", ",", "is", "working", "well", "?", "?"}},
}

func TestTokenize(t *testing.T) {
	for _, test := range valTokenize {
		got := Tokenize(test.sentence, true)
		for idx := range got {
			if got[idx] != test.want[idx] {
				t.Fatalf("=> Got: %v\n=> Want: %v", got, test.want)
			}
		}
	}
}
