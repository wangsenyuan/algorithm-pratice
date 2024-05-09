package main

import "testing"

func runSample(t *testing.T, s string, a []int, expect int) {
	res := solve(s, a)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "hhardh"
	a := []int{3, 2, 9, 11, 7, 1}
	expect := 5
	runSample(t, s, a, expect)
}

func TestSample2(t *testing.T) {
	s := "hhzarwde"
	a := []int{3, 2, 6, 9, 4, 8, 7, 1}
	expect := 4
	runSample(t, s, a, expect)
}

func TestSample3(t *testing.T) {
	s := "hhaarr"
	a := []int{1, 2, 3, 4, 5, 6}
	expect := 0
	runSample(t, s, a, expect)
}
