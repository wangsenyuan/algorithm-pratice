package main

import "testing"

func runSample(t *testing.T, a []int, l int, s int, expect int) {
	res := solve(a, l, s)
	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s, l := 2, 2
	a := []int{1, 3, 1, 2, 4, 1, 2}
	expect := 3
	runSample(t, a, l, s, expect)
}

func TestSample2(t *testing.T) {
	s, l := 2, 2
	a := []int{1, 100, 1, 100, 1, 100, 1}
	expect := -1
	runSample(t, a, l, s, expect)
}

func TestSample3(t *testing.T) {
	s, l := 565, 2
	a := []int{31, 76, 162, -182, -251, 214}
	expect := 1
	runSample(t, a, l, s, expect)
}
