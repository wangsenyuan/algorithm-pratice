package main

import "testing"

func runSample(t *testing.T, n int, X []int, Y []int, expect int) {
	res := solve(n, X, Y)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 3
	X := []int{0, 1, -4}
	Y := []int{0, 1, 5}
	expect := 3
	runSample(t, n, X, Y, expect)
}

func TestSample2(t *testing.T) {
	n := 3
	X := []int{0, 1, 3}
	Y := []int{1, 1, -1}
	expect := 2
	runSample(t, n, X, Y, expect)
}
