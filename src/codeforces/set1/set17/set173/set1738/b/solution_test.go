package main

import "testing"

func runSample(t *testing.T, n int, k int, b []int, expect bool) {
	res := solve(n, k, b)

	if res != expect {
		t.Fatalf("Sample expect %t, but got %t", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, k := 3, 3
	b := []int{2, 3, 4}
	expect := false
	runSample(t, n, k, b, expect)
}

func TestSample2(t *testing.T) {
	n, k := 3, 2
	b := []int{3, 4}
	expect := false
	runSample(t, n, k, b, expect)
}
