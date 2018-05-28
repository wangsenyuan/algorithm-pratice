package main

import "testing"

func runSample(t *testing.T, n int, A []int, expect int64) {
	res := solve(n, A)

	if res != expect {
		t.Errorf("Sample %d %v, expect %d, but got %d", n, A, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 6, []int{1, 2, 3, 4, 5, 6}, 5)
}

func TestSample2(t *testing.T) {
	runSample(t, 9, []int{0, 2, 1, 2, 1, 3, 0, 1, 0}, 31)
}

func TestSample3(t *testing.T) {
	runSample(t, 7, []int{0, 5, -5, 5, -5, 4, -4}, 20)
}
