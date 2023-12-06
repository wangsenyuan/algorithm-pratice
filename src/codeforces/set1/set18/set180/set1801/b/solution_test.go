package main

import "testing"

func runSample(t *testing.T, a []int, b []int, expect int) {
	res := solve(len(a), a, b)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{1, 2}
	b := []int{2, 1}
	expect := 0
	runSample(t, a, b, expect)
}

func TestSample2(t *testing.T) {
	a := []int{1, 2, 3, 4, 2}
	b := []int{5, 7, 3, 10, 5}
	expect := 1
	runSample(t, a, b, expect)
}

func TestSample3(t *testing.T) {
	a := []int{0, 1}
	b := []int{0, 1}
	expect := 1
	runSample(t, a, b, expect)
}
