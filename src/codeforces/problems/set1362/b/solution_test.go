package main

import "testing"

func runSample(t *testing.T, n int, A []int, expect int) {
	res := solve(n, A)

	if res != expect {
		t.Errorf("Sample %d %v, expect %d, but got %d", n, A, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 4
	A := []int{1, 0, 2, 3}
	expect := 1
	runSample(t, n, A, expect)
}

func TestSample2(t *testing.T) {
	n := 6
	A := []int{10, 7, 14, 8, 3, 12}
	expect := 4
	runSample(t, n, A, expect)
}

func TestSample3(t *testing.T) {
	n := 2
	A := []int{0, 1023}
	expect := 1023
	runSample(t, n, A, expect)
}

func TestSample4(t *testing.T) {
	n := 6
	A := []int{1, 4, 6, 10, 11, 12}
	expect := -1
	runSample(t, n, A, expect)
}
