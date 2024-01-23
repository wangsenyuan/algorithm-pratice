package main

import "testing"

func runSample(t *testing.T, a []int, expect int) {
	res := solve(a)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{1, 2}
	expect := 2
	runSample(t, a, expect)
}

func TestSample2(t *testing.T) {
	a := []int{2, 4, 5, 0, 1}
	expect := 9
	runSample(t, a, expect)
}

func TestSample3(t *testing.T) {
	a := []int{0, 4, 4, 4}
	expect := 0
	runSample(t, a, expect)
}
