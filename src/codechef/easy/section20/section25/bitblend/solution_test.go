package main

import "testing"

func runSample(t *testing.T, A []int, expect int) {
	res, _ := solve(A)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{0, 0}
	expect := -1
	runSample(t, A, expect)
}

func TestSample2(t *testing.T) {
	A := []int{6, 16, 69}
	expect := 1
	runSample(t, A, expect)
}
