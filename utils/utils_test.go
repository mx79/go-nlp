package utils

import (
	"testing"
)

func AssertPass(t *testing.T, got, want interface{}) {
	if got != want {
		t.Fatalf("Got %v, want %v", got, want)
	}
}

func AssertFail(t *testing.T, got, want interface{}) {
	if got == want {
		t.Fatalf("Didn't want %v", got)
	}
}

func TestListToStr(t *testing.T) {
	type listToStrTest struct {
		strSlice []string
		want     string
	}

	testsMustPass := []listToStrTest{
		{[]string{"I'm", "doing", "some", "tests"}, "I'm doing some tests"},
		{[]string{"I'm  ", "doing", "some", "tests"}, "I'm   doing some tests"},
	}

	testsMustFail := []listToStrTest{
		{[]string{"I'm", "doing ", "some ", "tests"}, "I'm doing some tests"},
		{[]string{"I'm", "doing", "some", "test"}, "I'm doing some tests"},
	}

	for _, test := range testsMustPass {
		got := ListToStr(test.strSlice)
		AssertPass(t, got, test.want)
	}

	for _, test := range testsMustFail {
		got := ListToStr(test.strSlice)
		AssertFail(t, got, test.want)
	}
}

func TestSliceContains(t *testing.T) {
	type sliceContainsTest[T Global] struct {
		slice []T
		value T
		got   bool
	}
}
