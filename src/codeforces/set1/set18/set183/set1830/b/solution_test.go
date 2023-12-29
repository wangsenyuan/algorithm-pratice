package main

import "testing"

func runSample(t *testing.T, a []int, b []int, expect int) {
	res := solve(a, b)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{2, 3, 2}
	b := []int{3, 3, 1}
	expect := 2
	runSample(t, a, b, expect)
}

func TestSample2(t *testing.T) {
	a := []int{4, 2, 8, 2, 1, 2, 7, 5}
	b := []int{3, 5, 8, 8, 1, 1, 6, 5}
	expect := 7
	runSample(t, a, b, expect)
}

func TestSample3(t *testing.T) {
	a := []int{4, 4, 8, 8, 8, 8, 8, 8}
	b := []int{8, 8, 8, 8, 8, 8, 8, 8}
	expect := 1
	runSample(t, a, b, expect)
}
