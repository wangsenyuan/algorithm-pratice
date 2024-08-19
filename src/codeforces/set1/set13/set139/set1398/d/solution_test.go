package main

import "testing"

func runSample(t *testing.T, r []int, b []int, g []int, expect int) {
	res := solve(r, b, g)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	r := []int{3}
	b := []int{5}
	g := []int{4}
	expect := 20
	runSample(t, r, b, g, expect)
}

func TestSample2(t *testing.T) {
	r := []int{9, 5}
	b := []int{1}
	g := []int{2, 8, 5}
	expect := 99
	runSample(t, r, b, g, expect)
}

func TestSample3(t *testing.T) {
	r := []int{11, 7, 20, 15, 19, 14, 2, 4, 13, 14}
	b := []int{8}
	g := []int{11}
	expect := 372
	runSample(t, r, b, g, expect)
}
