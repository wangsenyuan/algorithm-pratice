package main

import "testing"

func runSample(t *testing.T, a []int, expect int) {
	res := solve(a)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{0, 3, 2, 8, 12}
	expect := 6
	runSample(t, a, expect)
}

func TestSample2(t *testing.T) {
	a := []int{2, 2, 4, 6, 3, 100, 100, 25}
	expect := 7
	runSample(t, a, expect)
}
