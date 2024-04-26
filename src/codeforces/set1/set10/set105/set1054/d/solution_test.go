package main

import "testing"

func runSample(t *testing.T, k int, a []int, expect int) {
	res := solve(k, a)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	k := 2
	a := []int{1, 3, 0}
	expect := 5
	runSample(t, k, a, expect)
}

func TestSample2(t *testing.T) {
	k := 3
	a := []int{1, 4, 4, 7, 3, 4}
	expect := 19
	runSample(t, k, a, expect)
}
