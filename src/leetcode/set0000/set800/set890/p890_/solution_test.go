package p890_

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, words []string, pattern string, expect []string) {
	res := findAndReplacePattern(words, pattern)
	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample %v %s, expect %v, but got %v", words, pattern, expect, res)
	}
}

func TestSample1(t *testing.T) {
	words := []string{"abc", "deq", "mee", "aqq", "dkd", "ccc"}
	pattern := "abb"
	expect := []string{"mee", "aqq"}
	runSample(t, words, pattern, expect)
}
