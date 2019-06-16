package p918

import "testing"

func runSample(t *testing.T, A []int, expect int) {
	res := maxSubarraySumCircular(A)
	if res != expect {
		t.Errorf("Sample %v expect %d, but got %d", A, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, []int{1, -2, 3, -2}, 3)
}

func TestSample2(t *testing.T) {
	runSample(t, []int{5, -3, 5}, 10)
}

func TestSample3(t *testing.T) {
	runSample(t, []int{3, -1, 2, -1}, 4)
}

func TestSample4(t *testing.T) {
	runSample(t, []int{3, -2, 2, -3}, 3)
}

func TestSample5(t *testing.T) {
	runSample(t, []int{-2, -3, -1}, -1)
}
