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
	A := []int{1, 25}
	runSample(t, n, A, 1)
}

func TestSample2(t *testing.T) {
	n := 5
	A := []int{27, 3, 15, 1024, 15}
	expect := 6
	runSample(t, n, A, expect)
}
