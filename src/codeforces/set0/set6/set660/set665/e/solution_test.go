package main

import "testing"

func runSample(t *testing.T, k int, a []int, expect int) {
	res := solve(k, a)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	k := 1
	a := []int{1, 2, 3}
	expect := 5
	runSample(t, k, a, expect)
}

func TestSample2(t *testing.T) {
	k := 2
	a := []int{1, 2, 3}
	expect := 3
	runSample(t, k, a, expect)
}