package p1739

import "testing"

func runSample(t *testing.T, n int, expect int) {
	res := minimumBoxes(n)

	if res != expect {
		t.Errorf("Sample %d, expect %d, but got %d", n, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 1000000000, 1650467)
}
