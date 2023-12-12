package main

import "testing"

func runSample(t *testing.T, n int, a []int, expect int) {
	res := solve(n, a)
	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 3
	a := []int{1}
	expect := 1
	runSample(t, n, a, expect)
}

func TestSample2(t *testing.T) {
	n := 4
	a := []int{1, 4}
	expect := 2
	runSample(t, n, a, expect)
}

func TestSample3(t *testing.T) {
	n := 11
	a := []int{4, 8}
	expect := 6720
	runSample(t, n, a, expect)
}
