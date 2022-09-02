package main

import "testing"

func runSample(t *testing.T, A []int, x int, expect int) {
	res := solve(A, x)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{25, 3, 3, 17, 8, 6, 1, 16, 15, 25, 17, 23}
	x := 8
	expect := 2
	runSample(t, A, x, expect)
}
