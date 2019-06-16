package p748

import "testing"

func runSample(t *testing.T, lp string, words []string, expect string) {
	res := shortestCompletingWord(lp, words)
	if res != expect {
		t.Errorf("sample (%s, %v) should give %s, but got %s", lp, words, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, "1s3 PSt", []string{"step", "steps", "stripe", "stepple"}, "steps")
}

func TestSample2(t *testing.T) {
	runSample(t, "1s3 456", []string{"looks", "pest", "stew", "show"}, "pest")
}
