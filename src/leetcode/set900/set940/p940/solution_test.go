package p940

import "testing"

func runSample(t *testing.T, S string, expect int) {
	res := distinctSubseqII(S)

	if res != expect {
		t.Errorf("Sample %s, expect %d, but got %d", S, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, "abc", 7)
}

func TestSample2(t *testing.T) {
	runSample(t, "aba", 6)
}

func TestSample3(t *testing.T) {
	runSample(t, "aaa", 3)
}
