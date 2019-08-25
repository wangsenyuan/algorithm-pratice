package p1035

import "testing"

func runSample(t *testing.T, A []int, B []int, expect int) {
	res := maxUncrossedLines(A, B)
	if res != expect {
		t.Errorf("Sample %v %v, expect %d, but got %d", A, B, expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{1, 4, 2}
	B := []int{1, 2, 4}
	runSample(t, A, B, 2)
}

func TestSample2(t *testing.T) {
	A := []int{2, 5, 1, 2, 5}
	B := []int{10, 5, 2, 1, 5, 2}
	runSample(t, A, B, 3)
}

func TestSample3(t *testing.T) {
	A := []int{1, 3, 7, 1, 7, 5}
	B := []int{1, 9, 2, 5, 1}
	runSample(t, A, B, 2)
}
