package main

import "testing"

func runSample(t *testing.T, a []int, b []int, k int, expect int) {
	res := solve(a, b, k)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{100}
	b := []int{100}
	k := 2
	expect := 10000
	runSample(t, a, b, k, expect)
}
