package main

import "testing"

func runSample(t *testing.T, a []int, expect int) {
	res := solve(a)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{2, 1, 5, 6, 4, 3}
	expect := 8
	runSample(t, a, expect)
}

func TestSample2(t *testing.T) {
	a := []int{4, 4, 4, 4, 4}
	expect := 8
	runSample(t, a, expect)
}

func TestSample3(t *testing.T) {
	a := []int{1, 1000000000}
	expect := 1000000000
	runSample(t, a, expect)
}
