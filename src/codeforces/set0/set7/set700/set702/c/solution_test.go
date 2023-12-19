package main

import "testing"

func runSample(t *testing.T, a []int, b []int, expect int) {
	res := solve(a, b)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{-2, 2, 4}
	b := []int{-3, 0}
	expect := 4
	runSample(t, a, b, expect)
}

func TestSample2(t *testing.T) {
	a := []int{1, 5, 10, 14, 17}
	b := []int{4, 11, 15}
	expect := 3
	runSample(t, a, b, expect)
}
