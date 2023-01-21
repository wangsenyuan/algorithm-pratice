package main

import "testing"

func runSample(t *testing.T, A []int, expect int) {
	res := solve(len(A), A)

	if res != expect {
		t.Errorf("Sample %v, expect %d, but got %d", A, expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{1, 7, 1, 6, 4, 4, 6}
	expect := 3
	runSample(t, A, expect)
}

func TestSample2(t *testing.T) {
	A := []int{1, 1, 4, 6, 4, 6, 4, 7}
	expect := 2
	runSample(t, A, expect)
}

func TestSample3(t *testing.T) {
	A := []int{3, 3, 3}
	expect := 0
	runSample(t, A, expect)
}

func TestSample4(t *testing.T) {
	A := []int{2, 5, 2, 3, 1, 4}
	expect := 4
	runSample(t, A, expect)
}
