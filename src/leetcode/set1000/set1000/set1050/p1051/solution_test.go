package p1051

import "testing"

func runSample(t *testing.T, hs []int, expect int) {
	res := heightChecker(hs)
	if res != expect {
		t.Errorf("Sample %v, expect %d, but got %d", hs, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, []int{1, 1, 4, 2, 1, 3}, 3)
}
