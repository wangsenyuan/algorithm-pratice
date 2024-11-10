package main

import "testing"

func runSample(t *testing.T, a []int, expect int) {
	res := solve(a)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{3, -3, -2, 5, -4}
	expect := 4
	runSample(t, a, expect)
}

func TestSample2(t *testing.T) {
	a := []int{5, -2, -4, -6, 2, 3, -6, 5, 3, -2}
	expect := 9
	runSample(t, a, expect)
}

func TestSample3(t *testing.T) {
	a := []int{-1, -2, -3, -4}
	expect := 0
	runSample(t, a, expect)
}

func TestSample4(t *testing.T) {
	a := []int{-14, 0, -15, 10, 8, -7, 7, 10, 1}
	expect := 9
	runSample(t, a, expect)
}
