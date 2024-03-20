package main

import "testing"

func runSample(t *testing.T, k int, x int, a []int, expect int) {
	res := solve(a, k, x)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	k, x := 2, 3
	a := []int{5, 1, 3, 10, 1}
	expect := 18
	runSample(t, k, x, a, expect)
}

func TestSample2(t *testing.T) {
	k, x := 3, 1
	a := []int{1, 100, 1, 1}
	expect := 100
	runSample(t, k, x, a, expect)
}

func TestSample3(t *testing.T) {
	k, x := 1, 5
	a := []int{10, 30, 30, 70, 10, 10}
	expect := -1
	runSample(t, k, x, a, expect)
}
