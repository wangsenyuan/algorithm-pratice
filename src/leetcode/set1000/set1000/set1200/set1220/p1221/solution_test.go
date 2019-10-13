package p5222

import "testing"

func runSample(t *testing.T, s string, expect int) {
	res := balancedStringSplit(s)
	if res != expect {
		t.Errorf("Sample %s, expect %d, but got %d", s, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, "RLRRLLRLRL", 4)
}

func TestSample2(t *testing.T) {
	runSample(t, "RLLLLRRRLR", 3)
}