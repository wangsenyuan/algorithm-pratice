package main

import "testing"

func runSample(t *testing.T, n int, A []int, expect int64) {
	res := solve(n, A)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 5
	A := []int{1, 2, 3, 4, 5}
	var expect int64 = 4
	runSample(t, n, A, expect)
}

func TestSample2(t *testing.T) {
	n := 7
	A := []int{1, 2, 1, 2, 1, 2, 1}
	var expect int64 = 10
	runSample(t, n, A, expect)
}

func TestSample3(t *testing.T) {
	n := 8
	A := []int{1, 8, 2, 7, 3, 6, 4, 5}
	var expect int64 = 16
	runSample(t, n, A, expect)
}
