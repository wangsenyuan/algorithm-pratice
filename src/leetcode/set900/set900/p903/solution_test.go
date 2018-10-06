package p903

import "testing"

func runSample(t *testing.T, S string, expect int) {
	res := numPermsDISequence(S)
	if res != expect {
		t.Errorf("Sample %s, expect %d, but got %d", S, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, "DID", 5)
}

func TestSample2(t *testing.T) {
	runSample(t, "I", 1)
}
