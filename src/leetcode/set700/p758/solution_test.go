package p758

import "testing"

func runSample(t *testing.T, words []string, S string, expect string) {
	res := boldWords(words, S)
	if res != expect {
		t.Errorf("sample: %v %s, should give %s; but got %s", words, S, expect, res)
	}
}

func TestSample1(t *testing.T) {
	words := []string{"ab", "bc"}
	S := "aabcd"
	expect := "a<b>abc</b>d"
	runSample(t, words, S, expect)
}

func TestSample2(t *testing.T) {
	words := []string{"ab", "bc", "cd"}
	S := "aabcd"
	expect := "a<b>abcd</b>"
	runSample(t, words, S, expect)
}

func TestSample3(t *testing.T) {
	words := []string{"b", "dee", "a", "ee", "c"}
	S := "cebcecceab"
	expect := "<b>c</b>e<b>bc</b>e<b>cc</b>e<b>ab</b>"
	runSample(t, words, S, expect)
}
