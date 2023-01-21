package main

import "testing"

func runSample(t *testing.T, n int, r1 int, r2 int, r3 int, d int, A []int, expect int64) {
	res := solve(n, r1, r2, r3, d, A)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, r1, r2, r3, d := 4, 1, 3, 4, 3
	A := []int{3, 2, 5, 1}
	var expect int64 = 34
	runSample(t, n, r1, r2, r3, d, A, expect)
}

func TestSample2(t *testing.T) {
	n, r1, r2, r3, d := 4, 2, 4, 4, 1
	A := []int{4, 5, 1, 2}
	var expect int64 = 31
	runSample(t, n, r1, r2, r3, d, A, expect)
}
