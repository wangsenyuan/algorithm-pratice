package main

import "testing"

func runSample(t *testing.T, n int, a []int, expect int) {
	res := solve(n, a)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 11
	a := []int{4, 3, 4}
	expect := 2
	runSample(t, n, a, expect)
}

func TestSample2(t *testing.T) {
	n := 6
	a := []int{3, 2, 1}
	expect := 1
	runSample(t, n, a, expect)
}

func TestSample3(t *testing.T) {
	n := 4
	a := []int{1, 1, 1, 1}
	expect := 2
	runSample(t, n, a, expect)
}

func TestSample4(t *testing.T) {
	n := 5
	a := []int{1, 1, 1, 1, 1}
	expect := 2
	runSample(t, n, a, expect)
}
