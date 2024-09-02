package main

import "testing"

func runSample(t *testing.T, a []int, expect int) {
	res := solve(a)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{1, 1}
	expect := 1
	runSample(t, a, expect)
}

func TestSample2(t *testing.T) {
	a := []int{10, 24, 40, 80}
	expect := 40
	runSample(t, a, expect)
}

func TestSample3(t *testing.T) {
	a := []int{540, 648, 810, 648, 720, 540, 594, 864, 972, 648}
	expect := 54
	runSample(t, a, expect)
}
