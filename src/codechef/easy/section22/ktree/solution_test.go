package main

import "testing"

func runSample(t *testing.T, n int, K uint64, E [][]int, A []int, expect bool) {
	res := solve(n, K, E, A)

	if res != expect {
		t.Errorf("Sample %d %d %v %v, expect %t, but got %t", n, K, E, A, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 6
	var K uint64 = 85
	E := [][]int{
		{1, 2},
		{2, 3},
		{2, 4},
		{4, 5},
		{3, 6},
	}
	A := []int{3, 5, 4, 7, 1, 9}
	expect := true
	runSample(t, n, K, E, A, expect)
}
