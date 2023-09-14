package main

import "testing"

func runSample(t *testing.T, a []int, expect int) {
	res := solve(a)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{1, 2, 3}
	expect := 4
	runSample(t, a, expect)
}
