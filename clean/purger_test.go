package clean

import (
	"testing"
)

var (
	valFrPurgeText = []struct {
		sentence string
		want     string
	}{
		{"Je crois que je vais! faire, les tests maintenant?", "crois vais fair tests"},
		{"IL FAUT que çà mârche bien?? !OUI", "faut ca march bien oui"},
		{"on VA voir si ?ça marche là ou pas", "on va voir ca march"},
	}
	valEnPurgeText = []struct {
		sentence string
		want     string
	}{
		{"Let's see if IT works or not?", "see work"},
		{"I want to see if the test pass or NOT?", "want see test pass"},
		{",? !$$(( yes IT WORKS, text cleaned!!", "yes work text clean"},
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
