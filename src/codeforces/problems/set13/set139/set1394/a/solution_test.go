package main

import "testing"

func runSample(t *testing.T, n, d, m int, A []int, expect int64) {
	res := solve(n, d, m, A)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, d, m := 5, 2, 11
	A := []int{8, 10, 15, 23, 5}
	var expect int64 = 48
	runSample(t, n, d, m, A, expect)
}

func TestSample2(t *testing.T) {
	n, d, m := 20, 2, 16
	A := []int{20, 5, 8, 2, 18, 16, 2, 16, 16, 1, 5, 16, 2, 13, 6, 16, 4, 17, 21, 7}
	var expect int64 = 195
	runSample(t, n, d, m, A, expect)
}
