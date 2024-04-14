package main

import "testing"

func runSample(t *testing.T, a []int, expect int) {
	res := solve(a)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{2, 1, 2, 10, 2, 10, 10, 2, 2, 1, 10, 10, 10, 10, 1, 1, 10, 10}
	expect := 14
	runSample(t, a, expect)
}

func TestSample2(t *testing.T) {
	a := []int{6, 6, 6, 3, 6, 1000000000, 3, 3, 6, 6}
	expect := 9
	runSample(t, a, expect)
}
