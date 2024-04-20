package main

import "testing"

func runSample(t *testing.T, a []int, r int, expect int) {
	res := solve(a, r)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	r := 2
	a := []int{0, 1, 1, 0, 0, 1}
	expect := 3
	runSample(t, a, r, expect)
}

func TestSample2(t *testing.T) {
	r := 3
	a := []int{1, 0, 0, 0, 1}
	expect := 2
	runSample(t, a, r, expect)
}

func TestSample3(t *testing.T) {
	r := 3
	a := []int{0, 0, 1, 1, 0, 1, 0, 0, 0, 1}
	expect := 3
	runSample(t, a, r, expect)
}
