package main

import "testing"

func runSample(t *testing.T, n int, A []int, expect bool) {
	res := solve(n, A)

	if res != expect {
		t.Errorf("Sample expect %t, but got %t", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 5
	A := []int{5, 3, 2, 1, 4}
	expect := true
	runSample(t, n, A, expect)
}

func TestSample2(t *testing.T) {
	n := 2
	A := []int{2, 1}
	expect := false
	runSample(t, n, A, expect)
}
