package main

import "testing"

func runSample(t *testing.T, A []int, expect int) {
	res := solve(A)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{1, 2, 3, 5}
	expect := 0
	runSample(t, A, expect)
}

func TestSample2(t *testing.T) {
	A := []int{0}
	expect := -1
	runSample(t, A, expect)
}

func TestSample3(t *testing.T) {
	A := []int{5, 10, 2, 4, 0, 6, 1}
	expect := 1
	runSample(t, A, expect)
}

func TestSample4(t *testing.T) {
	A := []int{0, 1, 3, 5, 7}
	// k = 2,  0, 0, 1, 3, 5
	// k = 4,  0, 0, 0, 0, 1
	expect := 3
	runSample(t, A, expect)
}
