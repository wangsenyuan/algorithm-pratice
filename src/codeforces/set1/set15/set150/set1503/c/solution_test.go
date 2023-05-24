package main

import "testing"

func runSample(t *testing.T, n int, A []int, C []int, expect int64) {
	res := solve(n, A, C)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 3
	A := []int{1, 2, 4}
	C := []int{9, 1, 1}
	var expect int64 = 11
	runSample(t, n, A, C, expect)
}

func TestSample2(t *testing.T) {
	n := 6
	A := []int{4, 8, 3, 2, 7, 0}
	C := []int{2, 4, 0, 3, 1, 1}
	var expect int64 = 13
	runSample(t, n, A, C, expect)
}
