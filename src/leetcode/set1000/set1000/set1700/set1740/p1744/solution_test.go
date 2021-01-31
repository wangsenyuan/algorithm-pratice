package p1744

import "testing"

func runSample(t *testing.T, s string, expect bool) {
	res := checkPartitioning(s)

	if res != expect {
		t.Errorf("Sample %s, expect %t, but got %t", s, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, "abcbdd", true)
}

func TestSample2(t *testing.T) {
	runSample(t, "bcbddxy", false)
}
