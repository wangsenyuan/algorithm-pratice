package p1617

import "testing"

func runSample(t *testing.T, n int, expect int) {
	res := minimumOneBitOperations(n)

	if res != expect {
		t.Errorf("Sample %d, expect %d, but got %d", n, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 6, 4)
}

func TestSample2(t *testing.T) {
	runSample(t, 9, 14)
}

func TestSample3(t *testing.T) {
	runSample(t, 333, 393)
}
