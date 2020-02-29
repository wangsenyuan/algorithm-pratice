package main

import "testing"

func runSample(t *testing.T, n int, A []int, expect int) {
	res := solve(n, A)

	if res != expect {
		t.Errorf("Sample %d %v, expect %d, but got %d", n, A, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 7
	A := []int{375, 750, 723, 662, 647, 656, 619}
	expect := 2
	runSample(t, n, A, expect)
}
