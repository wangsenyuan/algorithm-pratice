package p1573

import "testing"

func runSample(t *testing.T, s string, expect int) {
	res := numWays(s)

	if res != expect {
		t.Errorf("Sample %s, expect %d, but got %d", s, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, "10101", 4)
}

func TestSample2(t *testing.T) {
	runSample(t, "0000", 3)
}

func TestSample3(t *testing.T) {
	runSample(t, "100100010100110", 12)
}
