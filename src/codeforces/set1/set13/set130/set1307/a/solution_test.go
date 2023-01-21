package main

import "testing"

func runSample(t *testing.T, n, d int, A []int, expect int) {
	res := solve(n, d, A)

	if res != expect {
		t.Errorf("Sample %d %d %v, expect %d, but got %d", n, d, A, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, d := 4, 5
	A := []int{1, 0, 3, 2}
	expect := 3
	runSample(t, n, d, A, expect)
}

func TestSample2(t *testing.T) {
	n, d := 2, 2
	A := []int{100, 1}
	expect := 101
	runSample(t, n, d, A, expect)
}

func TestSample3(t *testing.T) {
	n, d := 1, 8
	A := []int{0}
	expect := 0
	runSample(t, n, d, A, expect)
}
