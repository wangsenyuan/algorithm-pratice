package p1137

import "testing"

func runSample(t *testing.T, n int, expect int) {
	res := tribonacci(n)
	if res != expect {
		t.Errorf("Sample %d, expect %d, but got %d", n, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 25, 1389537)
}
