package main

import "testing"

func runSample(t *testing.T, a []int, expect int) {
	res := solve(a)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{3, 2, 3}
	expect := 7
	runSample(t, a, expect)
}

func TestSample2(t *testing.T) {
	a := []int{1, 2, 3, 4, 5}
	expect := 500000012
	runSample(t, a, expect)
}
