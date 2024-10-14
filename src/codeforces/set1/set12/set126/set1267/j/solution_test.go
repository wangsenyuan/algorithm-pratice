package main

import "testing"

func runSample(t *testing.T, n int, c []int, expect int) {
	res := solve(n, c)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 11
	c := []int{1, 5, 1, 5, 1, 5, 1, 1, 1, 1, 5}
	expect := 3
	runSample(t, n, c, expect)
}

func TestSample2(t *testing.T) {
	n := 6
	c := []int{1, 2, 2, 2, 2, 1}
	expect := 3
	runSample(t, n, c, expect)
}

func TestSample3(t *testing.T) {
	n := 5
	c := []int{4, 3, 3, 1, 2}
	expect := 4
	runSample(t, n, c, expect)
}
