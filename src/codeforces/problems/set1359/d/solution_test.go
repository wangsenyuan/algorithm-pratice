package main

import "testing"

func runSample(t *testing.T, n int, A []int, expect int) {
	res := solve(n, A)

	if res != expect {
		t.Errorf("Sample %d %v, expect %d, but got %d", n, A, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 5
	A := []int{5, -2, 10, -1, 4}
	expect := 6
	runSample(t, n, A, expect)
}

func TestSample2(t *testing.T) {
	n := 8
	A := []int{5, 2, 5, 3, -30, -30, 6, 9}
	expect := 10
	runSample(t, n, A, expect)
}

func TestSample3(t *testing.T) {
	n := 3
	A := []int{-10, 6, -15}
	expect := 0
	runSample(t, n, A, expect)
}
