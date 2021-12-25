package main

import "testing"

func runSample(t *testing.T, n int, C []int, H []int, expect bool) {
	res := solve(n, C, H)

	if res != expect {
		t.Errorf("Sample %d %v %v, expect %t, but got %t", n, C, H, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 5
	C := []int{1, 2, 3, 4, 5}
	H := []int{1, 2, 3, 4, 5}
	expect := false
	runSample(t, n, C, H, expect)
}

func TestSample2(t *testing.T) {
	n := 5
	C := []int{1, 2, 3, 4, 5}
	H := []int{5, 4, 3, 4, 5}
	expect := true
	runSample(t, n, C, H, expect)
}
