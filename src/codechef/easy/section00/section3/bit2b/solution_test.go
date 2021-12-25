package main

import "testing"

func runSample(t *testing.T, n int, A []int, x, expect int) {
	res := solve(n, A, x)
	if res != expect {
		t.Errorf("Sample %d %v %d, expect %d, but got %d", n, A, x, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 5
	A := []int{2, 8, 8, 2, 9}
	x := 2
	expect := 4
	runSample(t, n, A, x, expect)
}

func TestSample2(t *testing.T) {
	n := 5
	A := []int{9, 1, 1, 1, 1}
	x := 2
	expect := 1
	runSample(t, n, A, x, expect)
}

func TestSample3(t *testing.T) {
	n := 5
	A := []int{9, 1, 1, 1, 1, 1}
	x := 2
	expect := 1
	runSample(t, n, A, x, expect)
}

func TestSample4(t *testing.T) {
	n := 7
	A := []int{9, 1, 1, 1, 1, 1, 1}
	x := 2
	expect := 2
	runSample(t, n, A, x, expect)
}

func TestSample5(t *testing.T) {
	n := 5
	A := []int{2, 2, 1, 1, 1}
	x := 2
	expect := 3
	runSample(t, n, A, x, expect)
}
