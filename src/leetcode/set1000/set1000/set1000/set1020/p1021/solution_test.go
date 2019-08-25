package p1021

import "testing"

func runSample(t *testing.T, A []int, expect int) {
	res := maxScoreSightseeingPair(A)
	if res != expect {
		t.Errorf("Sample %v, expect %d, but got %d", A, expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{8, 1, 5, 2, 6}
	expect := 11
	runSample(t, A, expect)
}
