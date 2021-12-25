package main

import "testing"

func runSample(t *testing.T, n int, A []int, expect int) {
	res := solve(n, A)
	if res != expect {
		t.Errorf("Sample %d %v, expect %d, but got %d", n, A, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 3, []int{1, 2, 3}, 3)
}

func TestSample2(t *testing.T) {
	runSample(t, 5, []int{1, 3, 2, 2, 4}, 4)
}

func TestSample3(t *testing.T) {
	runSample(t, 5, []int{1, 3, 4, 2, 3}, 3)
}
