package main

import "testing"

func runSample(t *testing.T, n int, A []int, expect int64) {
	res := solve(n, A)

	if res != expect {
		t.Errorf("Sample %d %v, expect %d, but got %d", n, A, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 5, []int{2, 2, 4, 17, 8}, 4)
}

func TestSample2(t *testing.T) {
	runSample(t, 3, []int{4, 8, 16}, 3)
}

func TestSample3(t *testing.T) {
	runSample(t, 3, []int{6, 12, 36}, 6)
}
