package main

import "testing"

func runSample(t *testing.T, a []int, b []int, expect int) {
	res := solve(a, b)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{1}
	b := []int{3, 2}
	expect := 0
	runSample(t, a, b, expect)
}

func TestSample2(t *testing.T) {
	a := []int{5, 1, 5}
	b := []int{3, 8, 3, 3}
	expect := 1
	runSample(t, a, b, expect)
}

func TestSample3(t *testing.T) {
	a := []int{4, 3, 3, 2, 2, 1, 1}
	b := []int{1, 1, 1, 1, 3, 3, 3, 3}
	expect := 4
	runSample(t, a, b, expect)
}

func TestSample4(t *testing.T) {
	a := []int{9, 2, 8, 3, 7, 4, 6, 5}
	b := []int{1, 2, 3, 2, 1, 4, 5, 6, 5}
	expect := 4
	runSample(t, a, b, expect)
}
