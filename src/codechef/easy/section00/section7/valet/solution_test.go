package main

import "testing"

func runSample(t *testing.T, n int, T []int, X []int, Y []int, A []int, expect int64) {
	res := solve(n, T, X, Y, A)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 3
	T := []int{5, 8, 11}
	X := []int{3, 10, 4}
	Y := []int{1, 0, 1}
	A := []int{11, 13, 5}
	var expect int64 = 24
	runSample(t, n, T, X, Y, A, expect)
}
