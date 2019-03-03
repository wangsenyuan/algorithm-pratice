package p1003

import "testing"

func runSample(t *testing.T, S string, expect bool) {
	res := isValid(S)

	if res != expect {
		t.Errorf("Sample %s, expect %t, but got %t", S, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, "abcabcababcc", true)
}

func TestSample2(t *testing.T) {
	runSample(t, "abccba", false)
}

func TestSample3(t *testing.T) {
	runSample(t, "cababc", false)
}
