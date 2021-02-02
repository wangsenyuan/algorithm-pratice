package main

import "testing"

func runSample(t *testing.T, n int, A []int, expect int64) {
	res := solve(n, A)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 3
	A := []int{0, 1, 2}
	var expect int64 = 2
	runSample(t, n, A, expect)
}

func TestSample2(t *testing.T) {
	n := 4
	A := []int{1, 3, 1, 2}
	var expect int64 = 5
	runSample(t, n, A, expect)
}
