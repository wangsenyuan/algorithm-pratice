package main

import "testing"

func runSample(t *testing.T, n int, expect int) {
	res := solve(n)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	// a = [1, 2]
	// mex([1], 2) = 3
	// mex[[2], 2] = 3
	// mex[[1, 2], 3] = 5,  3, 4, 5
	runSample(t, 2, 12)
}

func TestSample2(t *testing.T) {
	// a = [1, 2]
	// mex([1], 2) = 3
	// mex[[2], 2] = 3
	// mex[[1, 2], 3] = 5,  3, 4, 5
	runSample(t, 3, 31)
}

func TestSample3(t *testing.T) {
	// a = [1, 2]
	// mex([1], 2) = 3
	// mex[[2], 2] = 3
	// mex[[1, 2], 3] = 5,  3, 4, 5
	runSample(t, 4999, 354226409)
}

func TestSample4(t *testing.T) {
	// a = [1, 2]
	// mex([1], 2) = 3
	// mex[[2], 2] = 3
	// mex[[1, 2], 3] = 5,  3, 4, 5
	runSample(t, 5, 184)
}

func TestSample5(t *testing.T) {
	// a = [1, 2]
	// mex([1], 2) = 3
	// mex[[2], 2] = 3
	// mex[[1, 2], 3] = 5,  3, 4, 5
	runSample(t, 1, 4)
}
