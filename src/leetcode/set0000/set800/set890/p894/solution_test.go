package p894

import "testing"

func runSample(t *testing.T, N int, expect int) {
	res := allPossibleFBT(N)
	if len(res) != expect {
		t.Errorf("Sample %d, expect %d, but got %d", N, expect, len(res))
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 7, 5)
}

func TestSample2(t *testing.T) {
	runSample(t, 19, 4892)
}
