package p792

import "testing"

func runSample(t *testing.T, s string, words []string, expect int) {
	res := numMatchingSubseq(s, words)
	if res != expect {
		t.Errorf("Sample %s %v, expect %d, but got %d", s, words, expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "abcde"
	words := []string{"a", "bb", "acd", "ace"}
	expect := 3
	runSample(t, s, words, expect)
}
