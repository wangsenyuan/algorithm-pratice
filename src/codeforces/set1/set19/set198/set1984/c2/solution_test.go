package main

import "testing"

func runSample(t *testing.T, a []int, expect int) {
	res := solve(a)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample(t *testing.T) {
	a := []int{2, -5, 3, -3}
	expect := 12
	runSample(t, a, expect)
}
