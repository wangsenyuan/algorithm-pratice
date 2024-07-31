package main

import "testing"

func runSample(t *testing.T, x int, a []int, expect int) {
	res := solve(a, x)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	x := 2
	a := []int{1, 1, 1, 1}
	expect := 8
	runSample(t, x, a, expect)
}

func TestSample2(t *testing.T) {
	x := 2
	a := []int{1, 2, 3}
	expect := 2
	runSample(t, x, a, expect)
}

func TestSample3(t *testing.T) {
	x := 6
	a := []int{10}
	expect := 0
	runSample(t, x, a, expect)
}

func TestSample4(t *testing.T) {
	x := 3
	a := []int{1, 2, 1, 4, 3, 8}
	expect := 10
	runSample(t, x, a, expect)
}
