package main

import "testing"

func runSample(t *testing.T, a []int, expect int) {
	res := solve(a)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{2, 2, 3, 4, 6}
	expect := 5
	runSample(t, a, expect)
}

func TestSample2(t *testing.T) {
	a := []int{2, 4, 6, 8}
	expect := -1
	runSample(t, a, expect)
}

func TestSample3(t *testing.T) {
	a := []int{2, 6, 9}
	expect := 4
	runSample(t, a, expect)
}
