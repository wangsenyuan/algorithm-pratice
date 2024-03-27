package main

import "testing"

func runSample(t *testing.T, a []int, k int, expect int) {
	res := solve(a, k)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	k := 2
	a := []int{2, 2, 2, 2}
	expect := 8
	runSample(t, a, k, expect)
}

func TestSample2(t *testing.T) {
	k := -3
	a := []int{3, -6, -3, 12}
	expect := 3
	runSample(t, a, k, expect)
}
