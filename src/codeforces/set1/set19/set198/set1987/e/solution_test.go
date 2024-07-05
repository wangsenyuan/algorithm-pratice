package main

import "testing"

func runSample(t *testing.T, n int, a []int, p []int, expect int) {
	res := solve(n, a, p)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{9, 3, 4, 1, 2}
	p := []int{1, 1, 3, 3}
	expect := 3
	runSample(t, len(a), a, p, expect)
}

func TestSample2(t *testing.T) {
	a := []int{5, 3}
	p := []int{1}
	expect := 2
	runSample(t, len(a), a, p, expect)
}
