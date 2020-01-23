package main

import "testing"

func runSample(t *testing.T, n int, A []int, expect int64) {
	res := solve(n, A)

	if res != expect {
		t.Errorf("Sample %d %v, expect %d, but got %d", n, A, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 10
	A := []int{7, 7, 5, 6, 4, 2, 3, 1, 1, 1}
	expect := int64(6)
	runSample(t, n, A, expect)
}
