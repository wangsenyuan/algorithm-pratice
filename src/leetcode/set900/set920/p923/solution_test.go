package p923

import "testing"

func runSample(t *testing.T, A []int, target int, expect int) {
	res := threeSumMulti(A, target)
	if res != expect {
		t.Errorf("Sample %v %d, expect %d, but got %d", A, target, expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{1, 1, 2, 2, 3, 3, 4, 4, 5, 5}
	target := 8
	expect := 20
	runSample(t, A, target, expect)
}

func TestSample2(t *testing.T) {
	A := []int{1, 1, 2, 2, 2, 2}
	target := 5
	expect := 12
	runSample(t, A, target, expect)
}

func TestSample3(t *testing.T) {
	A := []int{2, 2, 2, 2}
	target := 6
	expect := 4
	runSample(t, A, target, expect)
}
