package main

import "testing"

func runSample(t *testing.T, n, m int, X, Y [][]int, expect int) {
	res := solve(n, m, X, Y)

	if res != expect {
		t.Errorf("Sample %d %d %v %v, expect %d, but got %d", n, m, X, Y, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, m := 1, 1
	X := [][]int{{1, 2}}
	Y := [][]int{{2, 1}}
	expect := 1
	runSample(t, n, m, X, Y, expect)
}

func TestSample2(t *testing.T) {
	n, m := 1, 2
	X := [][]int{{1, 2}}
	Y := [][]int{{2, 1}, {1, 2}}
	expect := 1
	runSample(t, n, m, X, Y, expect)
}
