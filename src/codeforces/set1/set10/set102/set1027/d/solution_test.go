package main

import "testing"

func runSample(t *testing.T, c []int, a []int, expect int) {
	res := solve(c, a)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	c := []int{1, 2, 3, 2, 10}
	a := []int{1, 3, 4, 3, 3}
	expect := 3
	runSample(t, c, a, expect)
}

func TestSample2(t *testing.T) {
	c := []int{1, 1, 1, 1, 1, 1, 1}
	a := []int{2, 2, 2, 3, 6, 7, 6}
	expect := 2
	runSample(t, c, a, expect)
}
