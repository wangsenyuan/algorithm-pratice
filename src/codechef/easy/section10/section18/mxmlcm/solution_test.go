package main

import "testing"

func runSample(t *testing.T, n, m int, A []int, expect int) {
	res := solve(n, m, A)

	if res != expect {
		t.Errorf("Sample %d %d %v, expect %d, but got %d", n, m, A, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, m := 3, 2
	A := []int{2, 1, 2}
	runSample(t, n, m, A, 1)
}

func TestSample2(t *testing.T) {
	n, m := 4, 7
	A := []int{2, 5, 6, 3}
	runSample(t, n, m, A, 7)
}
