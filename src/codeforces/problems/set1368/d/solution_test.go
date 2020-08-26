package main

import "testing"

func runSample(t *testing.T, n int, A []int, expect int64) {
	res := solve(n, A)

	if res != expect {
		t.Errorf("Sample %d %v, expect %d, but got %d", n, A, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 1
	A := []int{123}
	var expect int64 = 15129

	runSample(t, n, A, expect)
}

func TestSample2(t *testing.T) {
	n := 3
	A := []int{1, 3, 5}
	var expect int64 = 51

	runSample(t, n, A, expect)
}

func TestSample3(t *testing.T) {
	n := 2
	A := []int{349525, 699050}
	var expect int64 = 1099509530625

	runSample(t, n, A, expect)
}
