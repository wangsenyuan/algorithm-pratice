package main

import "testing"

func runSample(t *testing.T, a []int, expect int) {
	res := solve(a)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{3, 1, 2}
	expect := 0
	runSample(t, a, expect)
}

func TestSample2(t *testing.T) {
	a := []int{4, 1, 5, 4, 1, 1}
	expect := 4
	runSample(t, a, expect)
}
