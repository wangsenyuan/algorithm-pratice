package main

import "testing"

func runSample(t *testing.T, a []int, b []int, p int, expect int) {
	res := solve(a, b, p)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{2, 3, 2, 1, 1, 3}
	b := []int{4, 3, 2, 6, 3, 6}
	expect := 16
	p := 3
	runSample(t, a, b, p, expect)
}
