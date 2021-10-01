package main

import "testing"

func runSample(t *testing.T, A []int, expect int) {
	res := solve(len(A), A)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{6, 5, 1, 2, 7, 8, 4, 3}
	expect := 0
	runSample(t, A, expect)
}

func TestSample2(t *testing.T) {
	A := []int{1, 6, 5, 2, 3, 4}
	expect := 1
	runSample(t, A, expect)
}

func TestSample3(t *testing.T) {
	A := []int{2, 5, 3, 3, 1, 3, 7}
	expect := 3
	runSample(t, A, expect)
}

func TestSample4(t *testing.T) {
	A := []int{6, 1, 4, 2, 4, 5}
	expect := 2
	runSample(t, A, expect)
}
