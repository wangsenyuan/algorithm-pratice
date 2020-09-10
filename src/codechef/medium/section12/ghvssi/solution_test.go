package main

import "testing"

func runSample(t *testing.T, n int, A []int, expect bool) {
	res := solve(n, A)

	if res != expect {
		t.Errorf("Sample %d %v, expect %t, but got %t", n, A, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 3
	A := []int{7, 9, 11}
	expect := true
	runSample(t, n, A, expect)
}

func TestSample2(t *testing.T) {
	n := 2
	A := []int{13, 15}
	expect := false
	runSample(t, n, A, expect)
}
