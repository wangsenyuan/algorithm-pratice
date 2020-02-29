package main

import "testing"

func runSample(t *testing.T, L, R uint64, A []int, expect uint64) {
	res := solve(L, R, A)

	if res != expect {
		t.Errorf("Sample %d %d %v, expect %d, but got %d", L, R, A, expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{5, 4, 3, 2, 1, 1, 2, 3, 4, 5}
	runSample(t, 21, 28, A, 6)
}

func TestSample2(t *testing.T) {
	A := []int{2, 3, 3, 3, 3, 2, 3, 3, 3, 3}
	runSample(t, 233, 23333, A, 19627)
}
