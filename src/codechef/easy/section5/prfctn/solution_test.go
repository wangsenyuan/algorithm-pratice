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
	A := []int{5, 9, 8, 3, 12}
	expect := 4
	runSample(t, n, A, expect)
}

func TestSample2(t *testing.T) {
	n := 2
	A := []int{19, 19}
	expect := 0
	runSample(t, n, A, expect)
}

func TestSample3(t *testing.T) {
	n := 5
	A := []int{1, 1, 2, 2, 3, 3, 3, 1}
	expect := 2
	runSample(t, n, A, expect)
}
