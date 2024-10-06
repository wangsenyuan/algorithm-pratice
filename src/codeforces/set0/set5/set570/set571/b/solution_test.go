package main

import "testing"

func runSample(t *testing.T, a []int, k int, expect int) {
	res := solve(a, k)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{1, 2, 4}
	k := 2
	expect := 1
	runSample(t, a, k, expect)
}

func TestSample2(t *testing.T) {
	a := []int{3, -5, 3, -5, 3}
	k := 2
	expect := 0
	runSample(t, a, k, expect)
}

func TestSample3(t *testing.T) {
	a := []int{4, 3, 4, 3, 2, 5}
	k := 3
	expect := 3
	runSample(t, a, k, expect)
}
