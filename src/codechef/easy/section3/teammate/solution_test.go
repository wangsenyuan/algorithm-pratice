package main

import "testing"

func runSample(t *testing.T, n int, A []int, expect int) {
	res := solve(n, A)

	if res != expect {
		t.Errorf("Sample %d %v, expect %d, but got %d", n, A, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 4, []int{1, 7, 3, 8}, 1)
}

func TestSample2(t *testing.T) {
	runSample(t, 4, []int{2, 3, 2, 2}, 3)
}
