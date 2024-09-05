package main

import "testing"

func runSample(t *testing.T, a []int, b []int, expect int) {
	res := solve(a, b)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{7, 2, 5}
	b := []int{15, 14, 3}
	expect := 6
	runSample(t, a, b, expect)
}
