package main

import "testing"

func runSample(t *testing.T, x int, c []int, expect int) {
	res := solve(x, c)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	x := 4
	c := []int{1, 1, 1}
	expect := 2
	runSample(t, x, c, expect)
}

func TestSample2(t *testing.T) {
	x := 6
	c := []int{1, 1, 1}
	expect := 3
	runSample(t, x, c, expect)
}
