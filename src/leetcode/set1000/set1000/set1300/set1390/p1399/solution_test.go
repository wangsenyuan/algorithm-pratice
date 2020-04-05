package p1399

import "testing"

func runSample(t *testing.T, n int, expect int) {
	res := countLargestGroup(n)

	if res != expect {
		t.Errorf("Sample %d expect %d, but got %d", n, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 13, 4)
}
