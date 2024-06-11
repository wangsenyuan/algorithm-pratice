package main

import "testing"

func runSample(t *testing.T, n int, s []int, l []int, expect int) {
	res := solve(n, s, l)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 2
	s := []int{1, 0, 1}
	l := []int{0, 1, 1}
	expect := 18
	runSample(t, n, s, l, expect)
}
