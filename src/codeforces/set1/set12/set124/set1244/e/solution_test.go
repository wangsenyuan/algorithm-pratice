package main

import "testing"

func runSample(t *testing.T, a []int, k int, expect int) {
	res := solve(a, k)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{3, 1, 7, 5}
	k := 5
	expect := 2
	runSample(t, a, k, expect)
}

func TestSample2(t *testing.T) {
	a := []int{4, 5, 5, 7, 5, 4, 5, 2, 4, 3}
	k := 9
	expect := 1
	runSample(t, a, k, expect)
}
