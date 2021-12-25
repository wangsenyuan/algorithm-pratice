package main

import "testing"

func runSample(t *testing.T, n int, A []int, expect int64) {
	res := solve(n, A)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 2
	A := []int{5, 9}
	var expect int64 = 12
	runSample(t, n, A, expect)
}

func TestSample2(t *testing.T) {
	n := 7
	A := []int{1, 2, 4, 8, 16, 64, 128}
	var expect int64 = 127
	runSample(t, n, A, expect)
}
