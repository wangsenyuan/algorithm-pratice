package main

import "testing"

func runSample(t *testing.T, a []int, b []int, expect int) {
	res := solve(a, b)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{1, 3, 6, 7}
	b := []int{9, 1, 1, 4}
	expect := 1

	runSample(t, a, b, expect)
}

func TestSample2(t *testing.T) {
	a := []int{1, 2, 3, 4, 5, 6, 7}
	b := []int{1, 1, 1, 1, 1, 1, 1}
	expect := 3

	runSample(t, a, b, expect)
}
