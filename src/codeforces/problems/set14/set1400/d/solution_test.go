package main

import "testing"

func runSample(t *testing.T, n int, A []int, expect int) {
	res := solve(n, A)

	if res != int64(expect) {
		t.Errorf("Sample %d %v, expect %d, but got %d", n, A, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 5
	A := []int{2, 2, 2, 2, 2}
	expect := 5
	runSample(t, n, A, expect)
}

func TestSample2(t *testing.T) {
	n := 6
	A := []int{1, 3, 3, 1, 2, 3}
	expect := 2
	runSample(t, n, A, expect)
}
