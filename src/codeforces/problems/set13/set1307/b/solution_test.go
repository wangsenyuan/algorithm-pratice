package main

import "testing"

func runSample(t *testing.T, n, x int, A []int, expect int) {
	res := solve(n, x, A)

	if res != expect {
		t.Errorf("Sample %d %d %v, expect %d, but got %d", n, x, A, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, x := 2, 4
	A := []int{1, 3}
	expect := 2
	runSample(t, n, x, A, expect)
}

func TestSample2(t *testing.T) {
	n, x := 3, 12
	A := []int{3, 4, 5}
	expect := 3
	runSample(t, n, x, A, expect)
}

func TestSample3(t *testing.T) {
	n, x := 1, 5
	A := []int{5}
	expect := 1
	runSample(t, n, x, A, expect)
}

func TestSample4(t *testing.T) {
	n, x := 2, 10
	A := []int{15, 4}
	expect := 2
	runSample(t, n, x, A, expect)
}
