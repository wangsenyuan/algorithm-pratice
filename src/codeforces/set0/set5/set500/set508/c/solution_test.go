package main

import "testing"

func runSample(t *testing.T, x int, r int, w []int, expect int) {
	res := solve(x, r, w)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	x, r := 8, 3
	w := []int{10}
	expect := 3
	runSample(t, x, r, w, expect)
}

func TestSample2(t *testing.T) {
	x, r := 12, 4
	w := []int{1, 2, 3, 4, 5, 7, 8, 9, 10}
	expect := 5
	runSample(t, x, r, w, expect)
}

func TestSample3(t *testing.T) {
	x, r := 1, 3
	w := []int{10}
	expect := -1
	runSample(t, x, r, w, expect)
}
