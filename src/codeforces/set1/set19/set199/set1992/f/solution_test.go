package main

import "testing"

func runSample(t *testing.T, a []int, x int, expect int) {
	res := solve(a, x)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{2, 3, 6, 2, 1, 2}
	x := 4
	expect := 3
	runSample(t, a, x, expect)
}

func TestSample2(t *testing.T) {
	a := []int{50000, 25000, 12500, 6250, 3125, 2, 4, 8, 16}
	x := 100000
	expect := 2
	runSample(t, a, x, expect)
}

func TestSample3(t *testing.T) {
	a := []int{1, 1, 1, 1, 1}
	x := 2
	expect := 1
	runSample(t, a, x, expect)
}
