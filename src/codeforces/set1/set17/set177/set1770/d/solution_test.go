package main

import "testing"

func runSample(t *testing.T, a []int, b []int, expect int) {
	res := solve(a, b)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{1, 2, 2}
	b := []int{1, 3, 3}
	expect := 6
	runSample(t, a, b, expect)
}

func TestSample2(t *testing.T) {
	a := []int{3, 3, 1, 3, 4}
	b := []int{4, 5, 2, 5, 5}
	expect := 0
	runSample(t, a, b, expect)
}
