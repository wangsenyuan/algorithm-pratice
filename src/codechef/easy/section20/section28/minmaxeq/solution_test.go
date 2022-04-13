package main

import "testing"

func runSample(t *testing.T, n, k int, A []int, expect int64) {
	res := solve(n, k, A)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, k := 3, 1
	A := []int{1, 2, 3}
	var expect int64 = 3
	runSample(t, n, k, A, expect)
}

func TestSample2(t *testing.T) {
	n, k := 4, 2
	A := []int{3, 3, 3, 1}
	var expect int64 = 4
	runSample(t, n, k, A, expect)
}

func TestSample3(t *testing.T) {
	n, k := 5, 1
	A := []int{1, 1, 1, 1, 1}
	var expect int64 = 7
	runSample(t, n, k, A, expect)
}
