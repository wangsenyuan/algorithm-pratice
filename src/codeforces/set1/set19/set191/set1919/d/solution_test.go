package main

import "testing"

func runSample(t *testing.T, a []int, expect bool) {
	res := solve(a)

	if res != expect {
		t.Fatalf("Sample expect %t, but got %t", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{2, 1, 0, 1, 1}
	expect := true
	runSample(t, a, expect)
}

func TestSample2(t *testing.T) {
	a := []int{1, 0, 2, 1, 3}
	expect := false
	runSample(t, a, expect)
}

func TestSample3(t *testing.T) {
	a := []int{12, 11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 0, 0, 1, 2, 3, 4, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	expect := false
	runSample(t, a, expect)
}
