package main

import "testing"

func runSample(t *testing.T, n int, a []int, expect int) {
	res := solve(n, a)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 10
	a := []int{1, 1, 32}
	expect := 2
	runSample(t, n, a, expect)
}

func TestSample2(t *testing.T) {
	n := 23
	a := []int{16, 1, 4, 1}
	expect := -1
	runSample(t, n, a, expect)
}

func TestSample3(t *testing.T) {
	n := 20
	a := []int{2, 1, 16, 1, 8}
	expect := 0
	runSample(t, n, a, expect)
}
