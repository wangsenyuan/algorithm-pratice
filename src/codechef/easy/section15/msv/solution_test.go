package main

import "testing"

func runSample(t *testing.T, n int, arr []int, expect int) {
	res := solve(n, arr)

	if res != expect {
		t.Errorf("Sample %d %v, expect %d, but got %d", n, arr, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 7
	arr := []int{8, 1, 28, 4, 2, 6, 7}
	expect := 3
	runSample(t, n, arr, expect)
}
