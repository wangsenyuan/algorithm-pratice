package main

import "testing"

func runSample(t *testing.T, a []int, expect bool) {
	res := solve(a)

	if res != expect {
		t.Fatalf("Sample expect %t, but got %t", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{5, 1, 3, 2, 1}
	expect := true
	runSample(t, a, expect)
}

func TestSample2(t *testing.T) {
	a := []int{7, 1, 5, 2, 1}
	expect := false
	runSample(t, a, expect)
}

func TestSample3(t *testing.T) {
	a := []int{1, 1, 4, 2, 1}
	expect := true
	runSample(t, a, expect)
}
