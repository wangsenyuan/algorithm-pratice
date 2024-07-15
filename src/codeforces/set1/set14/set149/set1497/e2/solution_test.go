package main

import "testing"

func runSample(t *testing.T, a []int, k int, expect int) {
	res := solve(a, k)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{18, 6, 2, 4, 1}
	k := 2
	expect := 1
	runSample(t, a, k, expect)
}

func TestSample2(t *testing.T) {
	a := []int{6, 8, 1, 24, 8}
	k := 0
	expect := 2
	runSample(t, a, k, expect)
}

func TestSample3(t *testing.T) {
	a := []int{6, 2, 2, 8, 9, 1, 3, 6, 3, 9, 7}
	k := 4
	expect := 2
	runSample(t, a, k, expect)
}
