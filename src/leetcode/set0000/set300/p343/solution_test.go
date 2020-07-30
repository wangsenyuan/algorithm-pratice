package p343

import "testing"

func runSample(t *testing.T, n int, expect int) {
	res := integerBreak(n)

	if res != expect {
		t.Errorf("Sample %d, expect %d, but got %d", n, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 10, 36)
}
