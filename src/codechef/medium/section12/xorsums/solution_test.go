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
	A := []int{7, 5, 13}
	expect := 8
	runSample(t, n, A, expect)
}

func TestSample2(t *testing.T) {
	n := 6
	A := []int{11, 4, 4, 13, 11, 5}
	expect := 14
	runSample(t, n, A, expect)
}
