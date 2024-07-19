package main

import "testing"

func runSample(t *testing.T, a []int, b []int, expect bool) {
	res := solve(a, b)

	if res != expect {
		t.Fatalf("Sample expect %t, but got %t", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{1, 2, 3, 4}
	b := []int{4, 3, 2, 1}
	expect := true
	runSample(t, a, b, expect)
}

func TestSample2(t *testing.T) {
	a := []int{1, 2, 3, 4}
	b := []int{1, 2, 3, 4}
	expect := true
	runSample(t, a, b, expect)
}

func TestSample3(t *testing.T) {
	a := []int{1, 5, 7, 1000, 4}
	b := []int{4, 1, 7, 5, 1000}
	expect := false
	runSample(t, a, b, expect)
}
