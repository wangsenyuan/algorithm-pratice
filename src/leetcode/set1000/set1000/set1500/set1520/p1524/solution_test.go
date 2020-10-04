package p1524

import "testing"

func runSample(t *testing.T, s string, expect int) {
	res := numSplits(s)

	if res != expect {
		t.Errorf("Sample %s, expect %d, but got %d", s, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, "aacaba", 2)
}
