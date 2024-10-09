package main

import "testing"

func runSample(t *testing.T, a []int, expect int) {
	res := solve(a)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{10, 9, 7, 8, 6, 5, 3, 4, 2, 1}
	expect := 2
	runSample(t, a, expect)
}

func TestSample2(t *testing.T) {
	a := []int{1, 2, 3, 4, 5, 6}
	expect := 0
	runSample(t, a, expect)
}

func TestSample3(t *testing.T) {
	a := []int{10, 7, 4, 2, 5, 8, 9, 6, 3, 1}
	expect := 4
	runSample(t, a, expect)
}
