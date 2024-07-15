package main

import "testing"

func runSample(t *testing.T, a []int, expect int) {
	res := solve(a, 0)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{18, 6, 2, 4, 1}
	expect := 3
	runSample(t, a, expect)
}

func TestSample2(t *testing.T) {
	a := []int{6, 8, 1, 24, 8}
	expect := 2
	runSample(t, a, expect)
}

func TestSample3(t *testing.T) {
	a := []int{1}
	expect := 1
	runSample(t, a, expect)
}
