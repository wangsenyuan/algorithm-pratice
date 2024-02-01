package main

import "testing"

func runSample(t *testing.T, a []int, expect int) {
	res := solve(a)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{2, 3}
	expect := 499122179
	runSample(t, a, expect)
}

func TestSample2(t *testing.T) {
	a := []int{0, 0, 0}
	expect := 1
	runSample(t, a, expect)
}

func TestSample3(t *testing.T) {
	a := []int{34124838, 0, 113193378, 8, 321939321, 113193378, 9463828, 99}
	expect := 77987843
	runSample(t, a, expect)
}
