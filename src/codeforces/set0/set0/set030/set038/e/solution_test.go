package main

import "testing"

func runSample(t *testing.T, x []int, c []int, expect int) {
	res := solve(x, c)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	x := []int{2, 3, 1}
	c := []int{3, 4, 2}
	expect := 5
	runSample(t, x, c, expect)
}

func TestSample2(t *testing.T) {
	x := []int{1, 3, 5, 6}
	c := []int{7, 1, 10, 1}
	expect := 11
	runSample(t, x, c, expect)
}
