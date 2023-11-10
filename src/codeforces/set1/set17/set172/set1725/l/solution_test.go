package main

import "testing"

func runSample(t *testing.T, a []int, expect int) {
	res := solve(a)
	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{2, -1, -1, 5, 2, -2, 9}
	expect := 4
	runSample(t, a, expect)
}

func TestSample2(t *testing.T) {
	a := []int{-1}
	expect := -1
	runSample(t, a, expect)
}
