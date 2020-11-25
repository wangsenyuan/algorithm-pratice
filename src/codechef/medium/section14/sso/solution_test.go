package main

import "testing"

func runSample(t *testing.T, n int, A []int, expect uint64) {
	res := solve(n, A)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 3
	A := []int{1, 9, 8}
	var expect uint64 = 27

	runSample(t, n, A, expect)
}

func TestSample2(t *testing.T) {
	n := 2
	A := []int{1, 1}
	var expect uint64 = 3

	runSample(t, n, A, expect)
}
