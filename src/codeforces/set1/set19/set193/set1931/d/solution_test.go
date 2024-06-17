package main

import "testing"

func runSample(t *testing.T, a []int, x int, y int, expect int) {
	res := solve(a, x, y)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	x, y := 5, 2
	a := []int{1, 2, 7, 4, 9, 6}
	expect := 2
	runSample(t, a, x, y, expect)
}

func TestSample2(t *testing.T) {
	x, y := 2, 3
	a := []int{14, 6, 1, 15, 12, 15, 8, 2, 15}
	expect := 7
	runSample(t, a, x, y, expect)
}
