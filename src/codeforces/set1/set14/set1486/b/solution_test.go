package main

import "testing"

func runSample(t *testing.T, X, Y []int, expect int64) {
	res := solve(len(X), X, Y)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	X := []int{0, 2}
	Y := []int{0, 0}
	runSample(t, X, Y, 3)
}
