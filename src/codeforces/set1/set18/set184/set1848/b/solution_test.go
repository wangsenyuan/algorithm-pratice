package main

import "testing"

func runSample(t *testing.T, k int, a []int, expect int) {
	res := solve(k, a)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	k := 2
	a := []int{1, 1, 2, 1, 1}
	expect := 0
	runSample(t, k, a, expect)
}
