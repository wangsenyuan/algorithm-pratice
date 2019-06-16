package p1017

import "testing"

func runSample(t *testing.T, N int, expect string) {
	res := baseNeg2(N)

	if res != expect {
		t.Errorf("Sample %d, expect %s, but got %s", N, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 2, "110")
}
