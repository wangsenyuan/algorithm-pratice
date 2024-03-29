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
	a := []int{2, 3, 1}
	expect := 6
	runSample(t, n, a, expect)
}

func TestSample2(t *testing.T) {
	n := 4
	a := []int{2, 1, 1, 1}
	expect := 8
	runSample(t, n, a, expect)
}

func TestSample3(t *testing.T) {
	n := 5
	a := []int{2, 4, 2, 5, 3}
	expect := 28
	runSample(t, n, a, expect)
}

func TestSample4(t *testing.T) {
	n := 4
	a := []int{2, 1, 4, 3}
	expect := 4
	runSample(t, n, a, expect)
}
