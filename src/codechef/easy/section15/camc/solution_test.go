package main

import "testing"

func runSample(t *testing.T, n, m int, A []int, expect int) {
	res := solve(n, m, A)

	if res != expect {
		t.Errorf("Sample %d %d %v, expect %d, but got %d", n, m, A, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, m := 4, 2
	A := []int{10, 20, 15, 28}
	expect := 5
	runSample(t, n, m, A, expect)
}
