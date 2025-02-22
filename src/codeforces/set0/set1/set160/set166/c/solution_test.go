package main

import "testing"

func runSample(t *testing.T, a []int, x int, expect int) {
	res := solve(a, x)
	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{10, 20, 30}
	x := 10
	expect := 1
	runSample(t, a, x, expect)
}

func TestSample2(t *testing.T) {
	a := []int{1, 2, 3}
	x := 4
	expect := 4
	runSample(t, a, x, expect)
}

func TestSample3(t *testing.T) {
	a := []int{1, 2, 3}
	x := 3
	expect := 2
	runSample(t, a, x, expect)
}
