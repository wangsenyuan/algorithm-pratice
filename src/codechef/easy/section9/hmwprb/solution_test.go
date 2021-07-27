package main

import "testing"

func runSample(t *testing.T, n int, k int, H []int, expect int) {
	res := solve(n, k, H)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, k := 7, 3
	H := []int{213, 124, 156, 263, 149, 6, 62}
	expect := 130
	runSample(t, n, k, H, expect)
}

func TestSample2(t *testing.T) {
	n, k := 7, 1
	H := []int{476, 457, 232, 130, 849, 95, 749}
	expect := 682
	runSample(t, n, k, H, expect)
}

func TestSample3(t *testing.T) {
	n, k := 7, 2
	H := []int{601, 35, 45, 304, 449, 452, 190}
	expect := 484
	runSample(t, n, k, H, expect)
}
