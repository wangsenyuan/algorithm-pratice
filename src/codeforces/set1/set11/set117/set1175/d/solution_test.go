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
	a := []int{-1, -2, 5, -4, 8}
	expect := 15
	runSample(t, a, k, expect)
}

func TestSample2(t *testing.T) {
	k := 6
	a := []int{-3, 0, -1, -2, -2, -4, -1}
	expect := -45
	runSample(t, a, k, expect)
}

func TestSample3(t *testing.T) {
	k := 1
	a := []int{3, -1, 6, 0}
	expect := 8
	runSample(t, a, k, expect)
}
