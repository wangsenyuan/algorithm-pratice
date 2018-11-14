package main

import "testing"

func runSample(t *testing.T, n int, A, B []int, expect int64) {
	res := solve(n, A, B)

	if res != expect {
		t.Errorf("Sample %d %v %v, expect %d, but got %d", n, A, B, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 3
	A := []int{4, 2, 2}
	B := []int{5, 3, 4}
	runSample(t, n, A, B, 3)
}

func TestSample2(t *testing.T) {
	n := 4
	A := []int{2, 1, 1, 2}
	B := []int{3, 3, 3, 3}
	runSample(t, n, A, B, 2)
}
