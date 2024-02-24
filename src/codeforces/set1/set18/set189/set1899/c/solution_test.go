package main

import "testing"

func runSample(t *testing.T, a []int, expect int) {
	res := solve(a)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{-10, 5, -8, 10, 6 - 10, 7, 9, -2, -6, 7, 2, -4, 6, -1, 7, -6, -7, 4, 1}
	expect := 10
	runSample(t, a, expect)
}

func TestSample2(t *testing.T) {
	a := []int{9, 9, 8, 8}
	expect := 17
	runSample(t, a, expect)
}

func TestSample3(t *testing.T) {
	a := []int{101, -99, 101}
	expect := 101
	runSample(t, a, expect)
}
