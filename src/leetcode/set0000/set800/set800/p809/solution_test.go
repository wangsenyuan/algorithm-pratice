package p809

import "testing"

func runSample(t *testing.T, S string, words []string, expect int) {
	res := expressiveWords(S, words)

	if res != expect {
		t.Errorf("sample %s %v, expect %d, but got %d", S, words, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, "heeellooo", []string{"hello", "hi", "helo"}, 1)
}
