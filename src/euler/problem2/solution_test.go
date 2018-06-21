package problem2

import "testing"

func runSample(t *testing.T, N int, expect int64) {
	res := EvenFibNumSum(N)
	if res != expect {
		t.Errorf("Sample %d, expect %d, but got %d", N, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 10, 10)
}

func TestSample2(t *testing.T) {
	runSample(t, 4000000, 4613732)
}
