package main

import "testing"

func runSample(t *testing.T, p []int, expect int) {
	res := solve(p)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	p := []int{-2, -2, 2, 2}
	// -2, -2, 0, -1
	expect := 5
	runSample(t, p, expect)
}

func TestSample2(t *testing.T) {
	p := []int{-3, -2, -1, 0, 1, 2, 3, 4}
	// -2, -2, 0, -1
	expect := 13
	runSample(t, p, expect)
}
