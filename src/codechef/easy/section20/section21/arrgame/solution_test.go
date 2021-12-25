package main

import "testing"

func runSample(t *testing.T, n int, A []int, expect bool) {
	res := solve(n, A)

	if res != expect {
		t.Errorf("Sample %d %v, expect %t, but got %t", n, A, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 7
	A := []int{1, 1, 0, 0, 0, 1, 1}
	expect := true

	runSample(t, n, A, expect)
}

func TestSample2(t *testing.T) {
	n := 8
	A := []int{1, 0, 1, 1, 1, 0, 0, 1}
	expect := false

	runSample(t, n, A, expect)
}

func TestSample3(t *testing.T) {
	n := 4
	A := []int{1, 1, 0, 1}
	expect := true

	runSample(t, n, A, expect)
}
