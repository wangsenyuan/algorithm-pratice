package p1968

import "testing"

func runSample(t *testing.T, p int, expect int) {
	res := minNonZeroProduct(p)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 60, 813987236)
}

func TestSample2(t *testing.T) {
	runSample(t, 40, 554966674)
}
