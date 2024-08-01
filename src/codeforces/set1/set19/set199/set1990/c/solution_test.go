package main

import "testing"

func runSample(t *testing.T, a []int, expect int) {
	res := solve(a)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{4, 4, 4, 4}
	expect := 40
	runSample(t, a, expect)
}

func TestSample2(t *testing.T) {
	a := []int{1, 1, 2, 2}
	// 0, 1, 1, 2
	// 0, 0, 1, 1
	// 0, 0, 0, 1
	expect := 13
	runSample(t, a, expect)
}

func TestSample3(t *testing.T) {
	a := []int{2, 1, 1, 2}
	// 0, 0, 1, 2
	// 0, 0, 0, 0
	expect := 9
	runSample(t, a, expect)
}
