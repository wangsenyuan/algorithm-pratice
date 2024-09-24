package main

import "testing"

func runSample(t *testing.T, a []int, k int, x int, expect int) {
	res := solve(a, k, x)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{12, 9}
	k, x := 1, 2
	expect := 30
	runSample(t, a, k, x, expect)
}
