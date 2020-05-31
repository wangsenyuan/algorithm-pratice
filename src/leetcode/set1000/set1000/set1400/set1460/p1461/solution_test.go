package p1461

import "testing"

func runSample(t *testing.T, s string, k int, expect bool) {
	res := hasAllCodes(s, k)

	if res != expect {
		t.Errorf("Sample %s %d, expect %t, but got %t", s, k, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, "00110", 2, true)
}

func TestSample2(t *testing.T) {
	runSample(t, "0110", 1, true)
}

func TestSample3(t *testing.T) {
	runSample(t, "0000000001011100", 4, false)
}
