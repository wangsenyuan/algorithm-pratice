package p873

import "testing"

func runSample(t *testing.T, A []int, expect int) {
	res := lenLongestFibSubseq(A)

	if res != expect {
		t.Errorf("Sample %v, expect %d, but got %d", A, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, []int{1, 2, 3, 4, 5, 6, 7, 8}, 5)
}

func TestSample2(t *testing.T) {
	runSample(t, []int{1, 3, 7, 11, 12, 14, 18}, 3)
}

func TestSample3(t *testing.T) {
	runSample(t, []int{1, 3, 5}, 0)
}
