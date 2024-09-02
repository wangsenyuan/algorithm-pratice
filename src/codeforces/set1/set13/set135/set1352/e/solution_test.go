package main

import "testing"

func runSample(t *testing.T, a []int, expect int) {
	res := solve(a)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{3, 1, 4, 1, 5, 9, 2, 6, 5}
	expect := 5
	runSample(t, a, expect)
}
