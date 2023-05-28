package main

import "testing"

func runSample(t *testing.T, n int, x []int, expect int) {
	res := solve(n, x)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 10
	x := []int{8, 7, 5, 4, 9, 4}
	expect := 3
	runSample(t, n, x, expect)
}

func TestSample2(t *testing.T) {
	n := 2
	x := []int{1, 1, 1, 1, 1, 1, 1, 1}
	expect := 1
	runSample(t, n, x, expect)
}

func TestSample3(t *testing.T) {
	n := 12
	x := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	expect := 4
	runSample(t, n, x, expect)
}
