package purger

import (
	"fmt"
	"github.com/mx79/go-nlp/clean"
	"testing"
)

var valPurgeText = []struct {
	name     string
	sentence string
	want     string
}{
	{},
}

func TestPurgeText(t *testing.T) {

}

var valRemovePunctuation = []struct {
	name     string
	sentence string
	want     string
}{
	{"valid", "dZLl! mep. :: ?dm, zm?", "dZLl mep  dm zm"},
	{"invalid", "dZLl! mep. :: ?dm, zm?", "dZLl! mep. :: ?dm, zm?"},
	{"valid", "!@#$%^&*()[]_+<>?:.,;", ""},
	{"invalid", "!@#$%^&*()[]_+<>?:.,;", "!@#$%^&*()[]_+<>?:.,;"},
	{"valid", "je fais, des ?tests !! est-ce que ça marche*", "je fais des tests est-ce que ça marche"},
	{"valid", "je fais, des ?tests !! est-ce que ça marche*", "je fais, des ?tests !! est-ce que ça marche*"},
}

func TestRemovePunctuation(t *testing.T) {
	for _, tt := range valRemovePunctuation {
		want := clean.RemovePunctuation(tt.sentence)
		if want != tt.want {
			fmt.Println("RemovePunct : échec du test")
		}
	}
}

var valRemoveAccent = []struct {
	name     string
	sentence string
	want     string
}{
	{},
}

func TestRemoveAccent(t *testing.T) {

}

var valStopwords = []struct {
	name     string
	sentence string
	want     string
}{
	{},
}

func TestStopword(t *testing.T) {

}

var valLemmatize = []struct {
	name     string
	sentence string
	want     string
}{
	{},
}

func TestLemmatize(t *testing.T) {

}

var valTokenize = []struct {
	name     string
	sentence string
	want     string
}{
	{},
}

func TestTokenize(t *testing.T) {

}

var valLower = []struct {
	name     string
	sentence string
	want     string
}{
	{},
}

func TestLower(t *testing.T) {

}

func main() {

	fmt.Println(clean.RemovePunctuation("dZLl! mep. :: ?dm, zm?"))
	fmt.Println(Stopword("Je fais des tests pas sûr du tout que ce que je suis en train de faire puisse fonctionner"))
	fmt.Println(clean.RemoveAccent("Est-ce que ça marche ce truc là"))
	fmt.Println(clean.Lower("La P dfgT nfD"))
	fmt.Println(clean.Tokenize("allo oui ça va monsieur")[0])
	fmt.Println(Lemmatize("Je fais des tests sur la lemmatization des mots avoir voudrais devoir"))
	fmt.Println(PurgeText("Là je vais? TEster tout:: en même temPS! pour que ça marche",
		true, true, true, true, true, false))
}
