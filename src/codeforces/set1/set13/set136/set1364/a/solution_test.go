package main

import "testing"

func runSample(t *testing.T, n, x int, A []int, expect int) {
	res := solve(n, x, A)
	if res != expect {
		t.Errorf("Sample %d %d %v, expect %d, but got %d", n, x, A, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, x := 3, 3
	A := []int{1, 2, 3}
	expect := 2
	runSample(t, n, x, A, expect)
}
