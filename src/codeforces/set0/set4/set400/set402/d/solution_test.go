package main

import "testing"

func runSample(t *testing.T, a []int, b []int, expect int) {
	res := solve(a, b)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{4, 20, 34, 10, 10}
	b := []int{2, 5}
	expect := -2
	runSample(t, a, b, expect)
}

func TestSample2(t *testing.T) {
	a := []int{2, 4, 8, 16}
	b := []int{3, 5, 7, 11, 17}
	expect := 10
	runSample(t, a, b, expect)
}
