package main

import "testing"

func runSample(t *testing.T, n int, x int, a []int, expect int) {
	res := solve(n, x, a)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, x := 3, 3
	a := []int{2, 3, 1}
	expect := 4
	runSample(t, n, x, a, expect)
}

func TestSample2(t *testing.T) {
	n, x := 7, 4
	a := []int{1, 3, 1, 2, 2, 4, 3}
	expect := 6
	runSample(t, n, x, a, expect)
}
