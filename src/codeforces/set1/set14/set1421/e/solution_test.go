package main

import "testing"

func runSample(t *testing.T, n int, A []int, expect int64) {
	res := solve(n, A)

	if res != expect {
		t.Errorf("Sample %d %v, expect %d, but got %d", n, A, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 4
	A := []int{5, 6, 7, 8}
	var expect int64 = 26
	runSample(t, n, A, expect)
}

func TestSample2(t *testing.T) {
	n := 5
	A := []int{4, -5, 9, -2, 1}
	var expect int64 = 15
	runSample(t, n, A, expect)
}

func TestSample3(t *testing.T) {
	n := 5
	A := []int{-32, 57, 44, -34, 66}
	var expect int64 = 233
	runSample(t, n, A, expect)
}
