package main

import "testing"

func runSample(t *testing.T, a []int, k int, expect int) {
	res := solve(a, k)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{1, 2, 2, 4, 2, 3}
	k := 5
	expect := 3
	runSample(t, a, k, expect)
}

func TestSample2(t *testing.T) {
	a := []int{3, 3, 2, 1, 1, 1, 3}
	k := 5
	expect := 4
	runSample(t, a, k, expect)
}
