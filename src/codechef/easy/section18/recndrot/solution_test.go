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
	A := []int{2, 1}
	expect := 2

	runSample(t, n, A, expect)
}

func TestSample2(t *testing.T) {
	n := 2
	A := []int{1, 2}
	expect := 1

	runSample(t, n, A, expect)
}

func TestSample3(t *testing.T) {
	n := 5
	A := []int{1, 3, 2, 1, 2}
	expect := 2

	runSample(t, n, A, expect)
}

func TestSample4(t *testing.T) {
	n := 7
	A := []int{7, 5, 3, 6, 4, 1, 2}
	expect := 4

	runSample(t, n, A, expect)
}

func TestSample5(t *testing.T) {
	n := 8
	A := []int{7, 5, 3, 2, 6, 4, 1, 2}
	expect := 4

	runSample(t, n, A, expect)
}
