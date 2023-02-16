package main

import "testing"

func runSample(t *testing.T, n int, A []int, expect int64) {
	res := solve(n, A)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 9
	A := []int{1, 2, 2, 2, 1, 1, 2, 2, 2}
	var expect int64 = 3
	runSample(t, n, A, expect)
}

func TestSample2(t *testing.T) {
	n := 12
	A := []int{1, 2, 3, 4, 3, 4, 2, 1, 3, 4, 2, 1}
	var expect int64 = 1
	runSample(t, n, A, expect)
}
