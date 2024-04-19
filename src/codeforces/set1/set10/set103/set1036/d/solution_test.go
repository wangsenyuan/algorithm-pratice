package main

import "testing"

func runSample(t *testing.T, a []int, b []int, expect int) {
	res := solve(a, b)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{11, 2, 3, 5, 7}
	b := []int{11, 7, 3, 7}
	expect := 3
	runSample(t, a, b, expect)
}

func TestSample2(t *testing.T) {
	a := []int{1, 2}
	b := []int{100}
	expect := -1
	runSample(t, a, b, expect)
}
