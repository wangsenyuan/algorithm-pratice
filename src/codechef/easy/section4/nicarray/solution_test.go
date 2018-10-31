package main

import "testing"

func runSample(t *testing.T, n int, S int, A []int, expect int) {
	res := solve(n, S, A)
	if res != expect {
		t.Errorf("Sample %d %d %v, expect %d, but got %d", n, S, A, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, s := 3, 3
	A := []int{1, 1, -1}
	expect := 3
	runSample(t, n, s, A, expect)
}

func TestSample2(t *testing.T) {
	n, s := 4, 8
	A := []int{1, -1, -1, 3}
	expect := 23
	runSample(t, n, s, A, expect)
}

func TestSample3(t *testing.T) {
	n, s := 3, 10
	A := []int{-1, -1, -1}
	expect := 150
	runSample(t, n, s, A, expect)
}
