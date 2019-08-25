package p1052

import "testing"

func runSample(t *testing.T, customes []int, grumpy []int, X int, expect int) {
	res := maxSatisfied(customes, grumpy, X)
	if res != expect {
		t.Errorf("Sample %v %v %d, expect %d, but got %d", customes, grumpy, X, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, []int{1, 0, 1, 2, 1, 1, 7, 5}, []int{0, 1, 0, 1, 0, 1, 0, 1}, 3, 16)
}
