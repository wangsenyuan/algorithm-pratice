package main

import "testing"

func runSample(t *testing.T, A []int, X int, expect int) {
	res := int(solve(A, X))

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{1, 2, 3, 1}
	X := 1
	expect := 10
	runSample(t, A, X, expect)
}

func TestSample2(t *testing.T) {
	A := []int{0, 0, 0}
	X := 21
	expect := 9
	runSample(t, A, X, expect)
}
