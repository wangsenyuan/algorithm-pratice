package main

import "testing"

func runSample(t *testing.T, a []int, expect int) {
	res := solve(a)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSamaple1(t *testing.T) {
	a := []int{4, -4, 1, -3, 1, -3}
	expect := 5
	runSample(t, a, expect)
}
