package main

import "testing"

func runSample(t *testing.T, n int, k int64, A [][]int, expect int) {
	res := solve(n, A, k)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 4
	var k int64 = 100
	A := [][]int{
		{13, 2, 1, 16},
		{15, 24, 3, 3},
		{5, 17, 9, 8},
		{9, 6, 11, 32},
	}
	expect := 9
	runSample(t, n, k, A, expect)
}

func TestSample2(t *testing.T) {
	n := 4
	var k int64 = 40
	A := [][]int{
		{13, 2, 1, 16},
		{15, 24, 3, 3},
		{5, 17, 9, 8},
		{9, 6, 11, 32},
	}
	expect := 6
	runSample(t, n, k, A, expect)
}
