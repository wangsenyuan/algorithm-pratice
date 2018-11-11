package main

import "testing"

func runSample(t *testing.T, n int, H []int, C []int, expect int) {
	res := solve(n, H, C)

	if res != expect {
		t.Errorf("Sample %d %v %v, expect %d, but got %d", n, H, C, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 5
	H := []int{1, 2, 3, 4, 5}
	C := []int{3, 3, 3, 3, 3}
	expect := 1
	runSample(t, n, H, C, expect)
}

func TestSample2(t *testing.T) {
	n := 5
	H := []int{3, 5, 1, 2, 3}
	C := []int{1, 2, 3, 4, 3}
	expect := 2
	runSample(t, n, H, C, expect)
}

func TestSample3(t *testing.T) {
	n := 5
	H := []int{5, 4, 3, 2, 3}
	C := []int{1, 2, 3, 4, 5}
	expect := 3
	runSample(t, n, H, C, expect)
}
