package main

import "testing"

func runSample(t *testing.T, n, k int, A []int, expect int) {
	res := solve(n, k, A)

	if res != expect {
		t.Errorf("Sample %d %d %v, expect %d, but got %d", n, k, A, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, k := 4, 2
	A := []int{1, 2, 3, 4}
	expect := 1
	runSample(t, n, k, A, expect)
}

func TestSample2(t *testing.T) {
	n, k := 4, 3
	A := []int{1, 2, 3, 4}
	expect := 2
	runSample(t, n, k, A, expect)
}

func TestSample3(t *testing.T) {
	n, k := 5, 3
	A := []int{5, 3, 4, 2, 6}
	expect := 2
	runSample(t, n, k, A, expect)
}

func TestSample4(t *testing.T) {
	n, k := 6, 4
	A := []int{5, 3, 50, 2, 4, 5}
	expect := 3
	runSample(t, n, k, A, expect)
}

func TestSample5(t *testing.T) {
	n, k := 4, 3
	A := []int{4, 3, 2, 1}
	expect := 2
	runSample(t, n, k, A, expect)
}
