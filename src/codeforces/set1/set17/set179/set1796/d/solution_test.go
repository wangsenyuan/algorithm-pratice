package main

import "testing"

func runSample(t *testing.T, a []int, k int, x int, expect int) {
	res := solve(a, k, x)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{2, -1, 2, 3}
	k := 1
	x := 2
	expect := 5
	runSample(t, a, k, x, expect)
}

func TestSample2(t *testing.T) {
	a := []int{-1, 2}
	k := 2
	x := 3
	expect := 7
	runSample(t, a, k, x, expect)
}

func TestSample3(t *testing.T) {
	a := []int{3, 2, 4}
	k := 0
	x := 5
	expect := 0
	runSample(t, a, k, x, expect)
}

func TestSample4(t *testing.T) {
	a := []int{4, -1, 9, -3, 7, -8}
	k := 2
	x := -8
	expect := 44
	runSample(t, a, k, x, expect)
}
