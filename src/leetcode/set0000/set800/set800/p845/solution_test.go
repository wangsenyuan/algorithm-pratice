package p845

import "testing"

func runSample(t *testing.T, A []int, expect int) {
	res := longestMountain(A)
	if res != expect {
		t.Errorf("Sample %v, expect %d, but got %d", A, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, []int{2, 1, 4, 7, 3, 2, 5}, 5)
}

func TestSample2(t *testing.T) {
	runSample(t, []int{2, 2, 2}, 0)
}

func TestSample3(t *testing.T) {
	runSample(t, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, 0)
}
