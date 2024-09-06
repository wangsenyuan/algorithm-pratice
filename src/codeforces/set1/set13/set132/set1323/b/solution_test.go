package main

import "testing"

func runSample(t *testing.T, k int, a []int, b []int, expect int) {
	res := solve(a, b, k)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{1, 0, 1}
	b := []int{1, 1, 1}
	k := 2
	expect := 4
	runSample(t, k, a, b, expect)
}

func TestSample2(t *testing.T) {
	a := []int{1, 1, 1}
	b := []int{1, 1, 1, 1, 1}
	k := 4
	expect := 14
	runSample(t, k, a, b, expect)
}
