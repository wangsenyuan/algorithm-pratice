package main

import "testing"

func runSample(t *testing.T, a []int, k int, expect int) {
	res := solve(a, k)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{1, 1, 1, 1, 1}
	k := 1
	expect := 31
	runSample(t, a, k, expect)
}

func TestSample2(t *testing.T) {
	a := []int{0, 1, 2, 3}
	k := 0
	expect := 10
	runSample(t, a, k, expect)
}
