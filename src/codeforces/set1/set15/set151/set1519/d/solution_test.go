package main

import "testing"

func runSample(t *testing.T, a []int, b []int, expect int) {
	res := solve(a, b)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{2, 3, 2, 1, 3}
	b := []int{1, 3, 2, 4, 2}
	expect := 29
	runSample(t, a, b, expect)
}

func TestSample2(t *testing.T) {
	a := []int{1, 8, 7, 6, 3, 6}
	b := []int{5, 9, 6, 8, 8, 6}
	expect := 235
	runSample(t, a, b, expect)
}
