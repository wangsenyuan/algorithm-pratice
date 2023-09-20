package main

import "testing"

func runSample(t *testing.T, n int, P []int, expect int) {
	res := solve(n, P)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 6
	P := []int{1, 2, 1, 4, 2}
	expect := 4
	runSample(t, n, P, expect)
}

func TestSample2(t *testing.T) {
	n := 2
	P := []int{1}
	expect := 2
	runSample(t, n, P, expect)
}
