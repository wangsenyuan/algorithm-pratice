package main

import "testing"

func runSample(t *testing.T, n int, k int, a []int, expect bool) {
	res := solve(n, k, a)

	if res != expect {
		t.Fatalf("Sample expect %t, but got %t", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, k := 12, 2
	a := []int{1, 1, 1, 1, 1, 7}
	expect := false
	runSample(t, n, k, a, expect)
}

func TestSample2(t *testing.T) {
	n, k := 12, 2
	a := []int{2, 2, 2, 2, 2, 2}
	expect := true
	runSample(t, n, k, a, expect)
}

func TestSample3(t *testing.T) {
	n, k := 37084, 5
	a := []int{7417, 7417, 7417, 7417, 7416}
	expect := true
	runSample(t, n, k, a, expect)
}
