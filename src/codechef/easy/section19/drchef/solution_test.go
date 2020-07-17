package main

import "testing"

func runSample(t *testing.T, n int, A []int, x int, expect int) {
	res := solve(n, A, x)

	if res != expect {
		t.Errorf("Sample %d %v, expect %d, but got %d", n, A, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, x := 5, 5
	A := []int{1, 2, 3, 4, 5}
	expect := 5
	runSample(t, n, A, x, expect)
}

func TestSample2(t *testing.T) {
	n, x := 5, 1
	A := []int{40, 30, 20, 10, 50}
	expect := 9
	runSample(t, n, A, x, expect)
}

func TestSample3(t *testing.T) {
	n, x := 3, 10
	A := []int{20, 1, 110}
	expect := 6
	runSample(t, n, A, x, expect)
}
