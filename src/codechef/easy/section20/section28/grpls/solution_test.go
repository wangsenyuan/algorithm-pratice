package main

import "testing"

func runSample(t *testing.T, x int, H []int, A []int, expect int) {
	res := solve(x, H, A)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	x := 5
	H := []int{1, 1, 2}
	A := []int{3, 2, 1, 1, 2, 1, 3}
	expect := 6
	runSample(t, x, H, A, expect)
}
