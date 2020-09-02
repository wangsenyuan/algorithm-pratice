package main

import "testing"

func runSample(t *testing.T, n, m int, A []int, expect int) {
	res := solve(n, m, A)

	if res != expect {
		t.Errorf("Sample %d %v, expect %d, but got %d", n, A, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, m := 4, 5
	A := []int{2, 3, 4, 3}
	expect := 1
	runSample(t, n, m, A, expect)
}

func TestSample2(t *testing.T) {
	n, m := 4, 5
	A := []int{2, -1, 4, 3}
	expect := 1
	runSample(t, n, m, A, expect)
}

func TestSample3(t *testing.T) {
	n, m := 4, 5
	A := []int{1, -1, -1, 5}
	expect := -1
	runSample(t, n, m, A, expect)
}

func TestSample4(t *testing.T) {
	n, m := 5, 4
	A := []int{-1, -1, 3, -1, -1}
	expect := 1
	runSample(t, n, m, A, expect)
}

func TestSample5(t *testing.T) {
	n, m := 5, 5
	A := []int{-1, -1, 3, -1, -1}
	expect := 0
	runSample(t, n, m, A, expect)
}
