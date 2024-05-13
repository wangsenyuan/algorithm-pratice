package main

import "testing"

func runSample(t *testing.T, d int, a []int, expect int) {
	res := solve(d, a)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	d := 3
	a := []int{1, 2, 3, 4}
	expect := 4
	runSample(t, d, a, expect)
}

func TestSample2(t *testing.T) {
	d := 2
	a := []int{-3, -2, -1, 0}
	expect := 2
	runSample(t, d, a, expect)
}

func TestSample3(t *testing.T) {
	d := 19
	a := []int{1, 10, 20, 30, 50}
	expect := 1
	runSample(t, d, a, expect)
}
