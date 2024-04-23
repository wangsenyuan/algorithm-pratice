package main

import "testing"

func runSample(t *testing.T, a []int, b []int, x int, expect int) {
	res := solve(a, b, x)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{1, 2, 3}
	b := []int{1, 2, 3}
	x := 9
	expect := 4
	runSample(t, a, b, x, expect)
}

func TestSample2(t *testing.T) {
	a := []int{5, 4, 2, 4, 5}
	b := []int{2}
	x := 5
	expect := 1
	runSample(t, a, b, x, expect)
}
