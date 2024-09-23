package main

import "testing"

func runSample(t *testing.T, l int, a []int, k int, expect int) {
	res := solve(l, a, k)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	l, k := 5, 3
	a := []int{5, 9, 1}
	// b = [5, 9, 1, 5, 9]
	//
	expect := 10
	runSample(t, l, a, k, expect)
}

func TestSample2(t *testing.T) {
	l, k := 10, 3
	a := []int{1, 2, 3, 4, 5}
	expect := 25
	runSample(t, l, a, k, expect)
}
