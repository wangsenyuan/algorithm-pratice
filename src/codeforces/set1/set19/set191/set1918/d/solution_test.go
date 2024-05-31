package main

import "testing"

func runSample(t *testing.T, a []int, expect int) {
	res := solve(a)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{1, 4, 5, 3, 3, 2}
	expect := 7
	runSample(t, a, expect)
}

func TestSample2(t *testing.T) {
	a := []int{1, 2, 3, 4, 5}
	expect := 5
	runSample(t, a, expect)
}

func TestSample3(t *testing.T) {
	a := []int{4, 1, 6, 3, 10, 7}
	expect := 11
	runSample(t, a, expect)
}
