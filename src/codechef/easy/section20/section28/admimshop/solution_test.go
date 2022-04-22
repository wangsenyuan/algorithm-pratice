package main

import "testing"

func runSample(t *testing.T, A []int, x int, expect int) {
	res := solve(A, x)

	if int(res) != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{1, 2, 3, 4}
	x := 2
	expect := 4
	runSample(t, A, x, expect)
}
