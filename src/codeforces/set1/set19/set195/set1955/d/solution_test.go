package main

import "testing"

func runSample(t *testing.T, a []int, b []int, k int, expect int) {
	res := solve(a, b, k)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{4, 1, 2, 3, 4, 5, 6}
	b := []int{1, 2, 3, 4}
	k := 2
	expect := 4
	runSample(t, a, b, k, expect)
}

func TestSample2(t *testing.T) {
	a := []int{4, 1, 2, 3, 4, 5, 6}
	b := []int{1, 2, 3, 4}
	k := 3
	expect := 3
	runSample(t, a, b, k, expect)
}
