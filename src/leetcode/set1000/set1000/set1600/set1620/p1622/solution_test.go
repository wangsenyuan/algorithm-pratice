package p1622

import "testing"

func runSample(t *testing.T, n int, k int, expect int) {
	res := numberOfSets(n, k)

	if res != expect {
		t.Errorf("Sample %d %d, expect %d, but got %d", n, k, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 4, 2, 5)
}

func TestSample2(t *testing.T) {
	runSample(t, 3, 1, 3)
}
