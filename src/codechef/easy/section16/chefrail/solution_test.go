package main

import "testing"

func runSample(t *testing.T, n, m int, X []int, Y []int, expect int64) {
	res := solve(n, m, X, Y)

	if res != expect {
		t.Errorf("Sample %d %d %v %v, expect %d, but got %d", n, m, X, Y, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, m := 2, 2
	X := []int{-4, 2}
	Y := []int{2, -8}
	var expect int64 = 1
	runSample(t, n, m, X, Y, expect)
}

func TestSample2(t *testing.T) {
	n, m := 8, 4
	X := []int{1, 2, 3, 6, -1, -2, -3, -6}
	Y := []int{6, -6, 1, -1}
	var expect int64 = 8
	runSample(t, n, m, X, Y, expect)
}
