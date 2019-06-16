package p846

import "testing"

func runSample(t *testing.T, A []int, W int, expect bool) {
	res := isNStraightHand(A, W)
	if res != expect {
		t.Errorf("Sample %v %d, expect %t, but got %t", A, W, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, []int{1, 2, 3, 6, 2, 3, 4, 7, 8}, 3, true)
}
