package main

import "testing"

func runSample(t *testing.T, p []int, q []int, expect int) {
	res := solve(p, q)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{1, 3, 2}
	b := []int{2, 1, 3}
	expect := 2
	runSample(t, a, b, expect)
}

func TestSample2(t *testing.T) {
	a := []int{7, 3, 6, 2, 1, 5, 4}
	b := []int{6, 7, 2, 5, 3, 1, 4}
	expect := 16
	runSample(t, a, b, expect)
}

func TestSample3(t *testing.T) {
	a := []int{1, 2, 3, 4, 5, 6}
	b := []int{6, 5, 4, 3, 2, 1}
	expect := 11
	runSample(t, a, b, expect)
}
