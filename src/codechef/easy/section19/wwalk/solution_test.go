package main

import "testing"

func runSample(t *testing.T, n int, A []int, B []int, expect int64) {
	res := solve(n, A, B)

	if res != expect {
		t.Errorf("Sample %d %v %v, expect %d, but got %d", n, A, B, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 4
	A := []int{1, 3, 3, 4}
	B := []int{1, 2, 4, 4}
	var expect int64 = 5

	runSample(t, n, A, B, expect)
}
