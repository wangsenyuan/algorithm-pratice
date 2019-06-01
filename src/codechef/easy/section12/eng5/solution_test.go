package main

import "testing"

func runSample(t *testing.T, n, m int, A []int, expect int) {
	res := solve(n, m, A)
	if res != expect {
		t.Errorf("Sample %d %d %v, expect %d, but got %d", n, m, A, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, m := 6, 3
	A := []int{1, 2, 3, 1, 2, 3}
	runSample(t, n, m, A, 3)
}
