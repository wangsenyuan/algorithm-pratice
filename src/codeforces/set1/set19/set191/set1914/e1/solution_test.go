package main

import "testing"

func runSample(t *testing.T, a []int, b []int, expect int) {
	res := solve(a, b)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{4, 2, 1}
	b := []int{1, 2, 4}
	expect := 1
	runSample(t, a, b, expect)
}

func TestSample2(t *testing.T) {
	a := []int{1, 20, 1, 20}
	b := []int{100, 15, 10, 20}
	expect := -9
	runSample(t, a, b, expect)
}

func TestSample3(t *testing.T) {
	a := []int{1000000000, 1000000000, 1000000000, 1000000000, 1000000000}
	b := []int{1, 1, 1, 1, 1}
	expect := 2999999997
	runSample(t, a, b, expect)
}
