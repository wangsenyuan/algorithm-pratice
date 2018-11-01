package main

import "testing"

func runSample(t *testing.T, n int, A []int, expect int) {
	res := solve(n, A)
	if res != expect {
		t.Errorf("Sample %d %v, expect %d, but got %d", n, A, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 3
	A := []int{-1, -1, -1}
	expect := -2
	runSample(t, n, A, expect)
}

func TestSample2(t *testing.T) {
	n := 5
	A := []int{1, -1, -1, 4, 1}
	expect := 4
	runSample(t, n, A, expect)
}

func TestSample3(t *testing.T) {
	n := 4
	A := []int{4, 6, 7, -1}
	expect := -1
	runSample(t, n, A, expect)
}
