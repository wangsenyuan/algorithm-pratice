package main

import "testing"

func runSample(t *testing.T, n int, X []int, Y []int, expect int) {
	res := solve(n, X, Y)

	if res != expect {
		t.Errorf("Sample %d %v %v, expect %d, but got %d", n, X, Y, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 5
	X := []int{0, -1, 1, 0, 0}
	Y := []int{0, 0, 0, 1, -1}
	expect := 1
	runSample(t, n, X, Y, expect)
}

func TestSample2(t *testing.T) {
	n := 5
	X := []int{31, 30, 20, 25, 25}
	Y := []int{11, -41, 14, 18, 38}
	expect := 1
	runSample(t, n, X, Y, expect)
}

func TestSample3(t *testing.T) {
	n := 2
	X := []int{0, 1}
	Y := []int{0, 1}
	expect := 4
	runSample(t, n, X, Y, expect)
}
