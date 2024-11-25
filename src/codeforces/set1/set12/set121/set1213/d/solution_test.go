package main

import "testing"

func runSample(t *testing.T, a []int, k int, expect int) {
	res := solve(a, k)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	k := 3
	a := []int{1, 2, 2, 4, 5}
	expect := 1
	runSample(t, a, k, expect)
}

func TestSample2(t *testing.T) {
	k := 3
	a := []int{1, 2, 3, 4, 5}
	expect := 2
	runSample(t, a, k, expect)
}

func TestSample3(t *testing.T) {
	k := 3
	a := []int{1, 2, 3, 3, 3}
	expect := 0
	runSample(t, a, k, expect)
}
