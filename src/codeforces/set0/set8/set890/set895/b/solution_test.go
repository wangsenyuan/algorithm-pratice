package main

import "testing"

func runSample(t *testing.T, a []int, k int, x int, expect int) {
	res := solve(a, k, x)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	x, k := 2, 1
	a := []int{1, 3, 5, 7}
	expect := 3
	runSample(t, a, k, x, expect)
}

func TestSample2(t *testing.T) {
	x, k := 2, 0
	a := []int{5, 3, 1, 7}
	expect := 4
	runSample(t, a, k, x, expect)
}

func TestSample3(t *testing.T) {
	x, k := 3, 1
	a := []int{3, 3, 3, 3, 3}
	expect := 25
	runSample(t, a, k, x, expect)
}

func TestSample4(t *testing.T) {
	x, k := 5, 0
	a := []int{4, 4, 4}
	expect := 9
	runSample(t, a, k, x, expect)
}

func TestSample5(t *testing.T) {
	x, k := 5, 0
	a := []int{3, 4}
	expect := 3
	runSample(t, a, k, x, expect)
}
