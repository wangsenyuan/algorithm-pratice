package p1543

import "testing"

func runSample(t *testing.T, s string, expect string) {
	res := makeGood(s)

	if res != expect {
		t.Errorf("Sample %s, expect %s, but got %s", s, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, "leEeetcode", "leetcode")
}

func TestSample2(t *testing.T) {
	runSample(t, "abBAcC", "")
}
