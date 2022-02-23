package main

import "testing"

func runSample(t *testing.T, n int, m int, s int, A [][]int, expect int64) {
	res := solve(n, m, s, A)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, m, s := 3, 2, 2
	A := [][]int{
		{1, 2},
		{4, 5},
		{7, 8},
	}
	var expect int64 = 9
	runSample(t, n, m, s, A, expect)
}

func TestSample2(t *testing.T) {
	n, m, s := 2, 2, 0
	A := [][]int{
		{100, 99},
		{98, 97},
	}
	var expect int64 = 394
	runSample(t, n, m, s, A, expect)
}
