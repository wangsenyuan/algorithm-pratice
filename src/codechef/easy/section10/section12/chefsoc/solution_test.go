package main

import "testing"

func runSample(t *testing.T, n int, A []int, expect int) {
	res := solve(n, A)

	if res != expect {
		t.Errorf("Sample %d %v, expect %d, but got %d", n, A, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 4, []int{1, 1, 1, 1}, 4)
}

func TestSample2(t *testing.T) {
	runSample(t, 3, []int{2, 2, 2}, 5)
}

func TestSample3(t *testing.T) {
	runSample(t, 4, []int{1, 2, 1, 1}, 6)
}
