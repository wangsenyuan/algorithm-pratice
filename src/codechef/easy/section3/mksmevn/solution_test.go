package main

import "testing"

func runSample(t *testing.T, A []int, expect int) {
	res := solve(len(A), A)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{7, 3, 5, 1}
	expect := 0
	runSample(t, A, expect)
}

func TestSample2(t *testing.T) {
	A := []int{4, 2, 7, 8, 10}
	expect := 1
	runSample(t, A, expect)
}

func TestSample3(t *testing.T) {
	A := []int{1, 1, 1}
	expect := -1
	runSample(t, A, expect)
}
