package main

import "testing"

func runSample(t *testing.T, a []int, x []int, expect int) {
	res := solve(a, x)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{3, 7, 9, 7, 8}
	x := []int{5, 2, 5, 7, 5}
	expect := 6
	runSample(t, a, x, expect)
}
