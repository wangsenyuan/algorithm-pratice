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
	A := []int{1, 4, 6}
	expect := 1
	runSample(t, n, A, expect)
}

func TestSample2(t *testing.T) {
	n := 4
	A := []int{2, 3, 6, 6}
	expect := 2
	runSample(t, n, A, expect)
}

func TestSample3(t *testing.T) {
	n := 3
	A := []int{6, 15, 10}
	expect := 3
	runSample(t, n, A, expect)
}

func TestSample4(t *testing.T) {
	n := 4
	A := []int{2, 3, 5, 7}
	expect := -1
	runSample(t, n, A, expect)
}

func TestSample5(t *testing.T) {
	n := 2
	A := []int{313, 313}
	expect := 2
	runSample(t, n, A, expect)
}
