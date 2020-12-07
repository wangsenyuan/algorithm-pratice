package main

import "testing"

func runSample(t *testing.T, X []int, Y []int, expect int) {
	res := solve(X, Y)

	if res != expect {
		t.Errorf("Sample %v %v, expect %d, but got %d", X, Y, expect, res)
	}
}

func TestSample1(t *testing.T) {
	X := []int{2, 6, 4, 0}
	Y := []int{2, 4}
	expect := 2
	runSample(t, X, Y, expect)
}

func TestSample2(t *testing.T) {
	X := []int{1, 9, 1, 9, 8, 1, 0}
	Y := []int{1, 1, 4, 5, 1, 4}
	expect := 0
	runSample(t, X, Y, expect)
}

func TestSample3(t *testing.T) {
	X := []int{179, 261, 432, 162, 82, 43, 10, 38}
	Y := []int{379, 357, 202, 184, 197}
	expect := 147
	runSample(t, X, Y, expect)
}
