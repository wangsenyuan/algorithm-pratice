package main

import "testing"

func runSample(t *testing.T, a []int, expect int) {
	res := solve(a)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{1, 1, 2}
	expect := 11
	runSample(t, a, expect)
}

func TestSample2(t *testing.T) {
	a := []int{1, 3, 3, 7}
	expect := 76
	runSample(t, a, expect)
}
