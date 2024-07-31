package main

import "testing"

func runSample(t *testing.T, k []int, x []int, c []int, expect int) {
	res := solve(k, x, c)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	k := []int{1, 1, 1}
	x := []int{1, 1, 1}
	c := []int{1, 1, 1, 1, 1}
	expect := 7
	runSample(t, k, x, c, expect)
}

func TestSample2(t *testing.T) {
	k := []int{2, 1, 1}
	x := []int{5, 1, 1}
	c := []int{1, 2, 3, 3, 5}
	expect := 13
	runSample(t, k, x, c, expect)
}
