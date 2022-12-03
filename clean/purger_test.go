package clean

import (
	"testing"
)

var (
	valFrPurgeText = []struct {
		sentence string
		want     string
	}{
		{"Je crois que je vais! faire, les tests maintenant?", "je crois vais faire tests maintenant"},
		{"IL FAUT que çà mârche bien?? !OUI", "il faut ca marche bien oui"},
		{"on VA voir si ?ça marche là ou pas", "on va voir ca march"},
	}
	valEnPurgeText = []struct {
		sentence string
		want     string
	}{
		{"Let's see if IT works or not?", "let's see it work not"},
		{"I want to see if the test pass or NOT?", "i want see test pass not"},
		{",? !$$(( yes IT WORKS, text cleaned!!", "yes it works text cleaned"},
	}
)

func TestPurgeText(t *testing.T) {
	frPurger := NewTextPurger(FR, true, true, true, true, true)
	for _, test := range valFrPurgeText {
		got := frPurger.Purge(test.sentence)
		AssertPass(t, got, test.want)
	}

	enPurger := NewTextPurger(EN, true, true, true, true, true)
	for _, test := range valEnPurgeText {
		got := enPurger.Purge(test.sentence)
		AssertPass(t, got, test.want)
	}
}
