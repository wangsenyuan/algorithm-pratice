package main

import "testing"

func runSample(t *testing.T, A []int, expect int) {
	res := solve(A)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{0, 2, 5, 5, 4, 8}
	expect := 3
	runSample(t, A, expect)
}

func TestSample2(t *testing.T) {
	A := []int{2, 0, 2, 1, 0, 0}
	expect := 1
	runSample(t, A, expect)
}

func TestSample3(t *testing.T) {
	A := []int{7, 1, 3, 4, 2, 10, 3, 9, 6}
	expect := 3
	runSample(t, A, expect)
}
