package main

import "testing"

func runSample(t *testing.T, a []int, k int, expect int) {
	res := solve(a, k)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{3, 9, 9, 2}
	k := 3
	expect := 3
	runSample(t, a, k, expect)
}

func TestSample2(t *testing.T) {
	a := []int{7, 1, 2, 6, 8, 3, 4, 5}
	k := 2
	expect := 6
	runSample(t, a, k, expect)
}
