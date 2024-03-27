package main

import "testing"

func runSample(t *testing.T, s string, x string, a []int, expect int) {
	res := solve(s, x, a)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "ababcba"
	x := "abb"
	a := []int{5, 3, 4, 1, 7, 6, 2}
	expect := 3
	runSample(t, s, x, a, expect)
}

func TestSample2(t *testing.T) {
	s := "bbbabb"
	x := "bb"
	a := []int{1, 6, 3, 4, 2, 5}
	expect := 4
	runSample(t, s, x, a, expect)
}
