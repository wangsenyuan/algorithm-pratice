package main

import "testing"

func runSample(t *testing.T, a []int, x int, expect int) {
	res := solve(a, x)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	x := 4
	a := []int{5, -1, 3, 4, -1}
	expect := 5
	runSample(t, a, x, expect)
}

func TestSample2(t *testing.T) {
	x := 0
	a := []int{-1, 2, -3}
	expect := 4
	runSample(t, a, x, expect)
}

func TestSample3(t *testing.T) {
	x := -1
	a := []int{-2, 1, -2, 3}
	expect := 3
	runSample(t, a, x, expect)
}
