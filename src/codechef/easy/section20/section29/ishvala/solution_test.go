package main

import "testing"

func runSample(t *testing.T, N int, M int, S int, X []int, Y []int, expect int) {
	res := int(solve(N, M, S, X, Y))

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	N, M, S := 7, 8, 2
	X := []int{5}
	Y := []int{4, 8}
	expect := 6
	runSample(t, N, M, S, X, Y, expect)
}

func TestSample2(t *testing.T) {
	N, M, S := 6, 5, 1
	X := []int{1, 4}
	Y := []int{}
	expect := 20
	runSample(t, N, M, S, X, Y, expect)
}
