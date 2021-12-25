package main

import "testing"

func runSample(t *testing.T, n int, A []int, expect bool) {
	res := solve(n, A)

	if res != expect {
		t.Errorf("Sample %d %v, expect %t, but got %t", n, A, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 4
	A := []int{1, 1, 2, 3}
	expect := true
	runSample(t, n, A, expect)
}

func TestSample2(t *testing.T) {
	n := 4
	A := []int{2, 1, 3, 3}
	expect := false
	runSample(t, n, A, expect)
}

func TestSample3(t *testing.T) {
	n := 5
	A := []int{5, 4, 4, 3, 1}
	expect := false
	runSample(t, n, A, expect)
}

func TestSample4(t *testing.T) {
	n := 5
	A := []int{3, 2, 1, 1, 4}
	expect := true
	runSample(t, n, A, expect)
}
