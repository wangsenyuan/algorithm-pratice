package main

import "testing"

func runSample(t *testing.T, a []int, expect int) {
	res := solve(a)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{3, 4, 2, 3, 4, 2, 2}
	expect := 3
	runSample(t, a, expect)
}

func TestSample2(t *testing.T) {
	a := []int{20, 1, 14, 10, 2}
	expect := 0
	runSample(t, a, expect)
}

func TestSample3(t *testing.T) {
	a := []int{5, 5, 4, 4, 3, 5, 7, 6, 5, 4, 4, 6, 5}
	expect := 21
	runSample(t, a, expect)
}
