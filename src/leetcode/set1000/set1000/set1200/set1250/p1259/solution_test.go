package p1259

import "testing"

func runSample(t *testing.T, n int, expect int) {
	res := numberOfWays(n)
	if res != expect {
		t.Errorf("Sample %d, expect %d, but got %d", n, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 2, 1)
	runSample(t, 4, 2)
	runSample(t, 6, 5)
	runSample(t, 8, 14)
}
