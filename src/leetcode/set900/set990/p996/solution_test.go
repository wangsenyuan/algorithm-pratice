package p996

import "testing"

func runSample(t *testing.T, A []int, expect int) {
	res := numSquarefulPerms(A)
	if res != expect {
		t.Errorf("Sample %v, expect %d, but got %d", A, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, []int{1, 17, 8}, 2)
}

func TestSample2(t *testing.T) {
	runSample(t, []int{2, 2, 2}, 1)
}

func TestSample3(t *testing.T) {
	runSample(t, []int{0, 0, 0, 1, 1, 1}, 4)
}
