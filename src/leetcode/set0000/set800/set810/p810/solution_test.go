package p810

import "testing"

func runSample(t *testing.T, A []int, L, M int, expect int) {
	res := maxSumTwoNoOverlap(A, L, M)
	if res != expect {
		t.Errorf("Sample %v %d %d, expect %d, but got %d", A, L, M, expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{0, 6, 5, 2, 2, 5, 1, 9, 4}
	L, M := 2, 1
	runSample(t, A, L, M, 20)
}

func TestSample2(t *testing.T) {
	A := []int{3, 8, 1, 3, 2, 1, 8, 9, 0}
	L, M := 2, 3
	runSample(t, A, L, M, 29)
}

func TestSample3(t *testing.T) {
	A := []int{3, 8, 1, 8, 9, 0, 3, 2, 1}
	L, M := 2, 3
	runSample(t, A, L, M, 29)
}
