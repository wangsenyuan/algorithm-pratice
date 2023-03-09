package main

import "testing"

func runSample(t *testing.T, X []int, W []int, expect int) {
	res := solve(X, W)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	X := []int{1, 5}
	W := []int{2, 1}
	expect := 1
	runSample(t, X, W, expect)
}

func TestSample2(t *testing.T) {
	X := []int{1, 20, 34}
	W := []int{1, 3, 2}
	expect := 125000015
	runSample(t, X, W, expect)
}

func TestSample3(t *testing.T) {
	X := []int{23, 54, 56, 85, 109, 126}
	W := []int{6, 3, 2, 5, 1, 4}
	expect := 312500100
	runSample(t, X, W, expect)
}
