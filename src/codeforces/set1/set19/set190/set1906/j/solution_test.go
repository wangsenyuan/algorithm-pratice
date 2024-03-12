package main

import "testing"

func runSample(t *testing.T, a []int, expect int) {
	res := solve(a)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{1, 2, 3}
	expect := 3

	runSample(t, a, expect)
}

func TestSample2(t *testing.T) {
	a := []int{1, 3, 2, 4, 5}
	expect := 17

	runSample(t, a, expect)
}

func TestSample3(t *testing.T) {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	expect := 379394847

	runSample(t, a, expect)
}
