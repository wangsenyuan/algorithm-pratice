package main

import "testing"

func runSample(t *testing.T, n, d int, A []int, expect int) {
	res := solve(n, d, A)
	if res != expect {
		t.Errorf("Sample %d %d %v, expect %d, but got %d", n, d, A, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, d := 5, 2
	A := []int{3, 4, 5, 2, 1}
	expect := 3
	runSample(t, n, d, A, expect)
}

func TestSample2(t *testing.T) {
	n, d := 5, 2
	A := []int{4, 3, 2, 1, 5}
	expect := -1
	runSample(t, n, d, A, expect)
}

func TestSample3(t *testing.T) {
	n, d := 5, 2
	A := []int{5, 4, 3, 2, 1}
	expect := 4
	runSample(t, n, d, A, expect)
}
