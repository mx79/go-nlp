package utils

import (
	"testing"
)

func AssertPass(t *testing.T, got, want interface{}) {
	if got != want {
		t.Fatalf("=> Got: %v\n=> Want: %v", got, want)
	}
}

var valListToStr = []struct {
	strSlice []string
	want     string
}{
	{[]string{"I'm", "doing", "some", "tests"}, "I'm doing some tests"},
	{[]string{"I'm  ", "doing", "some", "tests"}, "I'm   doing some tests"},
}

func TestListToStr(t *testing.T) {
	for _, test := range valListToStr {
		got := ListToStr(test.strSlice)
		AssertPass(t, got, test.want)
	}
}
