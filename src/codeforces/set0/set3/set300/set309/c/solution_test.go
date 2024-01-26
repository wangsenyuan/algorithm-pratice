package main

import "testing"

func runSample(t *testing.T, a []int, b []int, expect int) {
	res := solve(a, b)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{8, 4, 3, 2, 2}
	b := []int{3, 2, 2}
	expect := 2
	runSample(t, a, b, expect)
}

func TestSample2(t *testing.T) {
	a := []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	b := []int{0, 0, 0, 0, 0, 0}
	expect := 6
	runSample(t, a, b, expect)
}

func TestSample3(t *testing.T) {
	a := []int{4, 4, 3, 3, 3}
	b := []int{0, 0, 0, 0, 0, 2, 0, 2, 0, 0}
	expect := 10
	runSample(t, a, b, expect)
}
