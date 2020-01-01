package main

import "testing"

func runSample(t *testing.T, n int, A []int, expect int) {
	res := solve(n, A)

	if res != expect {
		t.Errorf("Sample %d %v, expect %d, but got %d", n, A, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 2
	A := []int{1, 2}
	expect := 3
	runSample(t, n, A, expect)
}

func TestSample2(t *testing.T) {
	n := 3
	A := []int{1, 1, 2}
	expect := 4
	runSample(t, n, A, expect)
}

func TestSample3(t *testing.T) {
	n := 5
	A := []int{1, 2, 3, 1, 2}
	expect := 8
	runSample(t, n, A, expect)
}
