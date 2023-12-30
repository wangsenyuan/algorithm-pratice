package main

import "testing"

func runSample(t *testing.T, a []int, expect int) {

	res := solve(a)

	if expect != res {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{1, 3, 3, 3, 7}
	expect := 2
	runSample(t, a, expect)
}

func TestSample2(t *testing.T) {
	a := []int{4, 2}
	expect := 2
	runSample(t, a, expect)
}

func TestSample3(t *testing.T) {
	a := []int{1, 1, 1, 1}
	expect := 1
	runSample(t, a, expect)
}

func TestSample4(t *testing.T) {
	a := []int{5, 4, 2, 1, 0, 0, 4}
	expect := 3
	runSample(t, a, expect)
}
