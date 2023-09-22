package main

import "testing"

func runSample(t *testing.T, a []int, expect int) {
	res := solve(a)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{2, 2, 0}
	expect := 5
	runSample(t, a, expect)
}

func TestSample2(t *testing.T) {
	a := []int{12, 11, 8, 8, 6, 6, 6, 5, 3, 2, 1}
	expect := 2596
	runSample(t, a, expect)
}
