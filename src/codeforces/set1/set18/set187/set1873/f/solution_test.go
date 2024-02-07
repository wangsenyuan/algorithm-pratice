package main

import "testing"

func runSample(t *testing.T, a []int, h []int, k int, expect int) {
	res := solve(a, h, k)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{3, 2, 4, 1, 8}
	h := []int{4, 4, 2, 4, 1}
	k := 12
	expect := 3
	runSample(t, a, h, k, expect)
}
