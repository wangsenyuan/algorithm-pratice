package main

import "testing"

func runSample(t *testing.T, a []int, k int, expect int) {
	res := solve(a, k)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	k := 2
	a := []int{2, 5, 2, 5}
	expect := 12
	runSample(t, a, k, expect)
}

func TestSample2(t *testing.T) {
	k := 4
	a := []int{9, 1, 8, 2, 7, 5, 6, 4}
	expect := 32
	runSample(t, a, k, expect)
}
