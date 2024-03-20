package main

import "testing"

func runSample(t *testing.T, a []int, expect int) {
	res := solve(a)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{0, 0, 1, 0}
	expect := 1
	runSample(t, a, expect)
}

func TestSample2(t *testing.T) {
	a := []int{1, 0, 1, 0, 1}
	expect := 3
	runSample(t, a, expect)
}
