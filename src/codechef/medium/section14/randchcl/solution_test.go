package main

import "testing"

func runSample(t *testing.T, n int, W []int, expect int) {
	res := solve(n, W)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 3
	W := []int{1, 2, 3}
	expect := 887328320
	runSample(t, n, W, expect)
}

func TestSample2(t *testing.T) {
	n := 4
	W := []int{2, 2, 2, 2}
	expect := 15
	runSample(t, n, W, expect)
}
