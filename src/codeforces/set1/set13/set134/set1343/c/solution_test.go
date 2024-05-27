package main

import "testing"

func runSample(t *testing.T, n int, A []int, expect int) {
	res := solve(n, A)

	if res != expect {
		t.Errorf("Sample %d %v, expect %d, but got %d", n, A, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 5, []int{1, 2, 3, -1, -2}, 2)
}

func TestSample2(t *testing.T) {
	runSample(t, 4, []int{-1, -2, -1, -3}, -1)
}

func TestSample3(t *testing.T) {
	runSample(t, 10, []int{-2, 8, 3, 8, -4, -15, 5, -2, -3, 1}, 6)
}

func TestSample4(t *testing.T) {
	runSample(t, 6, []int{1, -1000000000, 1, -1000000000, 1, -1000000000}, -2999999997)
}
