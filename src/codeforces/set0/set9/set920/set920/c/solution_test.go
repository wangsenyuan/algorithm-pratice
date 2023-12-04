package main

import "testing"

func runSample(t *testing.T, a []int, s string, expect bool) {
	res := solve(a, s)

	if res != expect {
		t.Fatalf("Sample expect %t, but got %t", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{1, 2, 5, 3, 4, 6}
	s := "01110"
	expect := true
	runSample(t, a, s, expect)
}

func TestSample2(t *testing.T) {
	a := []int{1, 2, 5, 3, 4, 6}
	s := "01010"
	expect := false
	runSample(t, a, s, expect)
}
