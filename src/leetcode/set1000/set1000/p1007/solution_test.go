package p1007

import "testing"

func runSample(t *testing.T, A, B []int, expect int) {
	res := minDominoRotations(A, B)
	if res != expect {
		t.Errorf("Sample %v %v, expect %d, but got %d", A, B, expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{2, 1, 2, 4, 2, 2}
	B := []int{5, 2, 6, 2, 3, 2}
	expect := 2
	runSample(t, A, B, expect)
}

func TestSample2(t *testing.T) {
	A := []int{3, 5, 1, 2, 3}
	B := []int{3, 6, 3, 3, 4}
	expect := -1
	runSample(t, A, B, expect)
}
